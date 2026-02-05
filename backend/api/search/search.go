package search

import (
	"context"
	"net/http"

	"github.com/Alia5/steaminputdb.com/api/search/configs"
	"github.com/Alia5/steaminputdb.com/api/search/games"
	"github.com/danielgtaylor/huma/v2"
)

func RegisterRoutes(a huma.API) {
	configs.RegisterRoutes(a)
	games.RegisterRoutes(a)
	registerRoute(a)
}

func registerRoute(a huma.API) {

	huma.Register(
		a,
		huma.Operation{
			Method:  http.MethodPost,
			Path:    "/v1/search/",
			Tags:    []string{"search"},
			Summary: "Search Steam games and configurations all at once",
			Description: `Search for games that are available on Steam
**and** Search for SteamInput configurations for Steam and non-Steam games
**all at once**

This endpoint supports no pagination and has pretty strict limits, it's intended for the "main-page" of the Frontend only
`,
			Errors: []int{
				422,
			},
		},
		Handler,
	)
}

func Handler(ctx context.Context, req *Request) (*Response, error) {
	gamesCh := make(chan *games.AppsResponse)
	configsCh := make(chan *configs.ConfigsResponse)
	errCh := make(chan error, 2)

	go func() {
		gamesResp, err := games.Handler(ctx, &games.Request{
			Body: games.AppsQueryBody{
				QueryText: req.Body.SearchTerm,
				Limit:     req.Body.LimitGames,
				Filter: games.AppFilter{
					IncludeGames: new(true),
					IncludeDemos: new(true),
					IncludeMods:  new(true),
				},
				Include: games.AppsInclude{
					Assets: true,
				},
			},
		})
		if err != nil {
			errCh <- err
			return
		}
		if gamesResp.Body == nil {
			errCh <- nil
			return
		}
		r, ok := gamesResp.Body.(*games.AppsResponse)
		if !ok {
			errCh <- nil
			return
		}
		gamesCh <- r
	}()

	go func() {
		configsResp, err := configs.Handler(ctx, &configs.Request{
			Body: configs.ConfigQueryBody{
				QueryText: req.Body.SearchTerm,
				ConfigPagination: configs.ConfigPagination{
					Limit: req.Body.LimitConfigs,
				},
			},
		})
		if err != nil {
			errCh <- err
			return
		}
		if configsResp.Body == nil {
			errCh <- nil
			return
		}
		r, ok := configsResp.Body.(*configs.ConfigsResponse)
		if !ok {
			errCh <- nil
			return
		}
		configsCh <- r
	}()

	var gamesResp *games.AppsResponse
	var configsResp *configs.ConfigsResponse

	for range 2 {
		select {
		case err := <-errCh:
			return nil, err
		case gamesResp = <-gamesCh:
		case configsResp = <-configsCh:
		}
	}

	return &Response{
		Body: SearchResponseBody{
			Games:   gamesResp.Items,
			Configs: configsResp.Items,
		},
	}, nil
}
