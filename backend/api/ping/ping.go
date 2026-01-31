package ping

import (
	"context"
	"net/http"

	"github.com/Alia5/steaminputdb.com/version"
	"github.com/danielgtaylor/huma/v2"
)

type Ping struct {
	Service string `json:"service"`
	Version string `json:"version"`
}

type PingResponse struct {
	Body Ping
}

func RegisterRoutes(a huma.API) {
	huma.Register(
		a,
		huma.Operation{
			Method:      http.MethodGet,
			Path:        "/v1/ping",
			Tags:        []string{"/v1"},
			Summary:     "/ping",
			Description: "Health check and basic information",
		},
		func(ctx context.Context, _ *struct{}) (*PingResponse, error) {
			return &PingResponse{
				Body: Ping{
					Service: "SteamInputDB.com",
					Version: version.Version,
				},
			}, nil
		},
	)
}
