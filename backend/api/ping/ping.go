package ping

import (
	"context"
	"net/http"

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
			Method: http.MethodGet,
			Path:   "/ping",
		},
		func(ctx context.Context, _ *struct{}) (*PingResponse, error) {
			return &PingResponse{
				Body: Ping{
					Service: "SteamInputDB.com",
					Version: "v1.0.0",
				},
			}, nil
		},
	)
}
