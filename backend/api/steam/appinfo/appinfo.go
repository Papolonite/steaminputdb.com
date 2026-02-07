package appinfo

import (
	"context"
	"errors"
	"net/http"

	"github.com/Alia5/steaminputdb.com/steamapi"
	"github.com/danielgtaylor/huma/v2"
)

type Response struct {
	Body *steamapi.AppInfo
}

type Request struct {
	AppID uint32 `query:"app_id" required:"true"`
}

func RegisterRoute(a huma.API) {
	huma.Register(
		a,
		huma.Operation{
			Method:      http.MethodGet,
			Path:        "/v1/steam/appinfo",
			Tags:        []string{"steam", "app"},
			Summary:     "Get Steam app info",
			Description: "Retrieve app information from Steam Store for a given app ID",
			Errors: []int{
				http.StatusBadGateway, http.StatusNotFound,
			},
		},
		func(c context.Context, req *Request) (*Response, error) {
			info, err := steamapi.GetStoreInfo(c, req.AppID)
			if err != nil {
				if errors.Is(err, steamapi.ErrInfoNotFound) {
					return nil, huma.Error404NotFound("app not found")
				}
				if errors.Is(err, steamapi.ErrRequest) {
					return nil, huma.Error502BadGateway("failed to get steam app info: %v", err)
				}
				return nil, err
			}

			return &Response{
				Body: info,
			}, nil
		},
	)
}
