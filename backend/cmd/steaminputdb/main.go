package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/Alia5/steaminputdb.com/api"
	"github.com/Alia5/steaminputdb.com/config"
	"github.com/Alia5/steaminputdb.com/logging"
	"github.com/Alia5/steaminputdb.com/metrics"
	"github.com/Alia5/steaminputdb.com/routes"
	"github.com/Alia5/steaminputdb.com/version"
	"github.com/alecthomas/kong"
	kongtoml "github.com/alecthomas/kong-toml"
	kongyaml "github.com/alecthomas/kong-yaml"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-fuego/fuego"
	"github.com/rs/cors"
)

const baseURL = "https://steaminputdb.com"

func main() {

	userCfg := findUserConfig(os.Args[1:])
	jsonPaths, yamlPaths, tomlPaths := configCandidatePaths(userCfg)

	var cfg config.Config
	_ = kong.Parse(&cfg,
		kong.Name("steaminputdb"),
		kong.Description(fmt.Sprintf("SteamInputDB Server - v%s", version.Version)),
		kong.UsageOnError(),
		// Load configuration from JSON/YAML/TOML in priority order; flags/env override config values.
		kong.Configuration(kong.JSON, jsonPaths...),
		kong.Configuration(kongyaml.Loader, yamlPaths...),
		kong.Configuration(kongtoml.Loader, tomlPaths...),
	)

	logging.SetupDefault(cfg.LogLevel)

	frontendListener, err := net.Listen("tcp", cfg.ListenAddress)
	if err != nil {
		panic(err)
	}
	apiListener, err := net.Listen("tcp", cfg.API.ListenAddress)
	if err != nil {
		panic(err)
	}
	metricsListener, err := net.Listen("tcp", cfg.Metrics.ListenAddress)
	if err != nil {
		panic(err)
	}

	frontendListener = addrDisplayListener{Listener: frontendListener, listenAddr: cfg.ListenAddress}
	apiListener = addrDisplayListener{Listener: apiListener, listenAddr: cfg.API.ListenAddress}
	metricsListener = addrDisplayListener{Listener: metricsListener, listenAddr: cfg.Metrics.ListenAddress}

	frontendSrv := fuego.NewServer(
		fuego.WithListener(frontendListener),
		fuego.WithEngineOptions(
			fuego.WithOpenAPIConfig(fuego.OpenAPIConfig{
				Disabled:         true,
				DisableSwaggerUI: true,
				DisableLocalSave: true,
			}),
		),
		fuego.WithLoggingMiddleware(
			fuego.LoggingConfig{
				DisableRequest:  true,
				DisableResponse: true,
			},
		),
		fuego.WithGlobalMiddlewares(
			metrics.Middleware,
			cors.New(cors.Options{
				AllowedOrigins:   []string{cfg.CorsOrigins},
				AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
				AllowedHeaders:   []string{"*"},
				AllowCredentials: true,
			}).Handler,
			logging.Middleware,
		),
	)

	apiSrv := fuego.NewServer(
		fuego.WithListener(apiListener),
		fuego.WithEngineOptions(
			fuego.WithErrorHandler(api.ErrorHandler),
			fuego.WithOpenAPIConfig(fuego.OpenAPIConfig{
				Disabled:         false,
				DisableSwaggerUI: false,
				DisableLocalSave: false,
				PrettyFormatJSON: true,
				UIHandler: func(specURL string) http.Handler {
					return http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
						w.Header().Set("Content-Type", "text/html; charset=utf-8")
						_, _ = w.Write([]byte(
							strings.Replace(
								fuego.DefaultOpenAPIHTML(specURL),
								"https://go-fuego.dev/img/logo.svg",
								fmt.Sprintf("%s/logo.svg", baseURL),
								2,
							),
						))
					})
				},
				Info: &openapi3.Info{
					Title:       "SteamInputDB API",
					Description: "API for SteamInputDB.com",
					License: &openapi3.License{
						Name: "AGPL-3.0",
						URL:  "https://www.gnu.org/licenses/agpl-3.0.en.html",
					},
					Version: version.Version,
				},
			}),
		),
		fuego.WithLoggingMiddleware(
			fuego.LoggingConfig{
				DisableRequest:  false,
				DisableResponse: false,
			},
		),
		fuego.WithGlobalMiddlewares(
			func(next http.Handler) http.Handler {
				return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					//revive:disable-next-line
					r = r.WithContext(context.WithValue(r.Context(), "requestURI", r.RequestURI))
					next.ServeHTTP(w, r)
				})
			},
			metrics.Middleware,
			cors.New(cors.Options{
				AllowedOrigins:   []string{cfg.API.CorsOrigins},
				AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
				AllowedHeaders:   []string{"*"},
				AllowCredentials: true,
			}).Handler,
			logging.Middleware,
		),
	)

	metricsSrv := fuego.NewServer(
		fuego.WithListener(metricsListener),
		fuego.WithEngineOptions(
			fuego.WithOpenAPIConfig(fuego.OpenAPIConfig{
				Disabled:         true,
				DisableSwaggerUI: true,
				DisableLocalSave: true,
			}),
		),
		fuego.WithLoggingMiddleware(
			fuego.LoggingConfig{
				DisableRequest:  true,
				DisableResponse: true,
			},
		),
		fuego.WithGlobalMiddlewares(
			metrics.Middleware,
			logging.Middleware,
		),
	)

	routes.Register(frontendSrv, apiSrv, metricsSrv)

	if cfg.API.PublicAddress != "" {
		apiSrv.OpenAPI.Description().Servers = []*openapi3.Server{
			{
				URL:         cfg.API.PublicAddress,
				Description: "Public API",
			},
		}
	}

	errChan := make(chan error, 3)
	var wg sync.WaitGroup

	servers := []*fuego.Server{frontendSrv, apiSrv, metricsSrv}
	for _, srv := range servers {
		wg.Add(1)
		go func(s *fuego.Server) {
			defer wg.Done()
			if err := s.Run(); err != nil {
				errChan <- err
			}
		}(srv)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-sigChan:
		slog.Info("shutdown signal received, shutting down servers")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := frontendSrv.Shutdown(ctx); err != nil {
			slog.Error("error shutting down frontend server", "err", err)
		}
		if err := apiSrv.Shutdown(ctx); err != nil {
			slog.Error("error shutting down API server", "err", err)
		}
		if err := metricsSrv.Shutdown(ctx); err != nil {
			slog.Error("error shutting down metrics server", "err", err)
		}

	case err := <-errChan:
		slog.Error("server error", "err", err)
		os.Exit(1)
	}

	wg.Wait()
}

type addrDisplayListener struct {
	net.Listener
	listenAddr string
}

func (l addrDisplayListener) Addr() net.Addr  { return l }
func (l addrDisplayListener) Network() string { return "tcp" }
func (l addrDisplayListener) String() string {
	addr := strings.TrimSpace(l.listenAddr)
	if addr == "" {
		return l.Listener.Addr().String()
	}

	host, port, err := net.SplitHostPort(addr)
	if err != nil {
		return l.Listener.Addr().String()
	}
	if host == "" || host == "0.0.0.0" || host == "::" {
		host = "localhost"
	}
	return net.JoinHostPort(host, port)
}

func findUserConfig(args []string) string {
	for i, a := range args {
		if strings.HasPrefix(a, "--config=") {
			return a[len("--config="):]
		}
		if a == "--config" && i+1 < len(args) {
			return args[i+1]
		}
	}
	return os.Getenv("CONFIG")
}

func configCandidatePaths(userPath string) (jsonPaths, yamlPaths, tomlPaths []string) {
	add := func(slice *[]string, p string) { *slice = append(*slice, p) }

	if userPath != "" {
		switch ext := filepath.Ext(userPath); ext {
		case ".json":
			add(&jsonPaths, userPath)
		case ".yaml", ".yml":
			add(&yamlPaths, userPath)
		case ".toml":
			add(&tomlPaths, userPath)
		default:
			add(&jsonPaths, userPath)
		}
	}

	wd, _ := os.Getwd()
	for _, base := range []string{"github.com/Alia5/steaminputdb.com", "steaminputdb", "config"} {
		add(&jsonPaths, filepath.Join(wd, base+".json"))
		add(&yamlPaths, filepath.Join(wd, base+".yaml"))
		add(&yamlPaths, filepath.Join(wd, base+".yml"))
		add(&tomlPaths, filepath.Join(wd, base+".toml"))
	}

	return
}
