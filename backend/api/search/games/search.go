package games

import (
	"context"
	"log/slog"
	"net/http"
	"reflect"
	"time"

	"github.com/Alia5/steaminputdb.com/steamapi"
	"github.com/danielgtaylor/huma/v2"
)

func RegisterRoutes(a huma.API) {
	registry := a.OpenAPI().Components.Schemas

	huma.Register(
		a,
		huma.Operation{
			Method:      http.MethodPost,
			Path:        "/v1/search/games",
			Tags:        []string{"search"},
			Summary:     "Search Steam games",
			Description: "Search for games that are available on Steam",
			Errors: []int{
				422,
			},
			Responses: map[string]*huma.Response{
				"200": {
					Content: map[string]*huma.MediaType{
						"application/json": {
							Schema: &huma.Schema{
								OneOf: []*huma.Schema{
									registry.Schema(reflect.TypeFor[AppsResponse](), true, ""),
									registry.Schema(reflect.TypeFor[steamapi.CStoreQuery_SearchSuggestions_Response](), true, ""),
								},
							},
						},
					},
				},
			},
		},
		Handler,
	)
}

func Handler(ctx context.Context, req *Request) (*Response, error) {

	query := &steamapi.CStoreQuery_SearchSuggestions_Request{
		SearchTerm: &req.Body.QueryText,
		MaxResults: &req.Body.Limit,
		Context: &steamapi.StoreBrowseContext{
			Language:    new("english"),
			CountryCode: new("US"),
		},
		Filters: &steamapi.CStoreQueryFilters{
			ReleasedOnly: new(true),
			TypeFilters: &steamapi.CStoreQueryFilters_TypeFilters{
				IncludeGames: req.Body.Filter.IncludeGames,
				IncludeDemos: req.Body.Filter.IncludeDemos,
				IncludeMods:  req.Body.Filter.IncludeMods,
			},
		},
		DataRequest: &steamapi.StoreBrowseItemDataRequest{
			IncludeAssets:    &req.Body.Include.Assets,
			IncludeBasicInfo: &req.Body.Include.Info,
			IncludeLinks:     &req.Body.Include.Links,
			IncludeRatings:   &req.Body.Include.Ratings,
			IncludePlatforms: &req.Body.Include.Platforms,
			IncludeRelease:   &req.Body.Include.Release,
		},
	}
	slog.Debug("Querying Steam SearchStore API with",
		"query", query,
	)

	queryResp, err := steamapi.DefaultClient.SearchSuggestions(ctx, query)
	if err != nil {
		return nil, huma.Error502BadGateway("Error querying Steam API", err)
	}

	if req.Body.Raw {
		return &Response{
			Body: (*raw)(queryResp),
		}, nil
	}

	items := make([]AppsSearchItem, len(queryResp.StoreItems))
	for i, storeItem := range queryResp.StoreItems {
		items[i].AppID = storeItem.Appid
		items[i].Name = storeItem.Name
		items[i].StoreUrlPath = storeItem.StoreUrlPath

		if storeItem.Type != nil {
			items[i].Type = TypeToString(storeItem.Type)
		}

		if storeItem.Links != nil {
			items[i].Links = &[]string{}
			for _, link := range storeItem.Links {
				if link == nil || link.Url == nil {
					continue
				}
				*items[i].Links = append(*items[i].Links, *link.Url)
			}
		}
		items[i].BasicInfo = storeItem.BasicInfo
		items[i].Assets = storeItem.Assets
		if storeItem.Platforms != nil {
			items[i].Platforms = AppsPlatforms{
				Windows:      storeItem.Platforms.Windows,
				SteamOSLinux: storeItem.Platforms.SteamosLinux,
				Mac:          storeItem.Platforms.Mac,
			}
		}

		if storeItem.Release != nil {
			items[i].Release = AppsRelease{}
			if storeItem.Release.SteamReleaseDate != nil {
				items[i].Release.SteamReleaseDate = time.Unix(int64(*storeItem.Release.SteamReleaseDate), 0)
			}
			if storeItem.Release.OriginalReleaseDate != nil && *storeItem.Release.OriginalReleaseDate != 0 {
				items[i].Release.OriginalReleaseDate = time.Unix(int64(*storeItem.Release.OriginalReleaseDate), 0)
			}
		}
	}

	return &Response{
		Body: &AppsResponse{
			Total: queryResp.Metadata.TotalMatchingRecords,
			Items: items,
		},
	}, nil
}

func TypeToString(t *int32) string {
	if t == nil {
		return ""
	}
	switch *t {
	case 0:
		return "game"
	case 1:
		return "demo"
	case 2:
		return "mod"
	default:
		return "unknown"
	}
}
