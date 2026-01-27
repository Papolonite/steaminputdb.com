package logging

import (
	"log/slog"
	"net/http"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)

		ctx := r.Context()
		statusCode := ctx.Value("status_code")
		duration := ctx.Value("duration")

		slog.Debug("request",
			"status_code", statusCode,
			"method", r.Method,
			"path", r.URL.Path,
			"duration_ms", duration,
			"request_id", w.Header().Get("X-Request-ID"),
			"remote_addr", r.RemoteAddr,
		)
	})
}
