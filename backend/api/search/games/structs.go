package games

import (
	"time"

	"github.com/Alia5/steaminputdb.com/steamapi"
)

type AppsInclude struct {
	Assets    bool `json:"assets,omitempty,omitzero" default:"false" doc:"Include app assets in the response"`
	Info      bool `json:"info,omitempty,omitzero" default:"false" doc:"Include app information in the response"`
	Links     bool `json:"links,omitempty,omitzero" default:"false" doc:"Include app links in the response"`
	Ratings   bool `json:"ratings,omitempty,omitzero" default:"false" doc:"Include app ratings in the response"`
	Platforms bool `json:"platforms,omitempty,omitzero" default:"false" doc:"Include app platform information in the response"`
	Release   bool `json:"release,omitempty,omitzero" default:"false" doc:"Include app release information in the response"`
}

type AppFilter struct {
	IncludeGames *bool `json:"include_games,omitempty" default:"true" doc:"Include games in the search results"`
	IncludeDemos *bool `json:"include_demos,omitempty" default:"false" doc:"Include demos in the search results"`
	IncludeMods  *bool `json:"include_mods,omitempty" default:"false" doc:"Include mods in the search results"`
}

type AppsQueryBody struct {
	QueryText string      `json:"query_text" required:"true" maxLength:"100" doc:"The search query string"`
	Limit     uint32      `json:"limit,omitempty,omitzero" default:"5" max:"50" doc:"Maximum number of results to return"`
	Filter    AppFilter   `json:"filter,omitzero"`
	Include   AppsInclude `json:"include,omitzero"`
	Raw       bool        `json:"raw,omitempty,omitzero" default:"false" doc:"Return raw Steam API response"`
}

type Request struct {
	Body AppsQueryBody `additionalProperties:"false"`
}

type AppsResponse struct {
	Total *int32           `json:"total" example:"123" doc:"Total number of matching items"`
	Items []AppsSearchItem `json:"items"`
}

type AppsSearchItem struct {
	AppID        *uint32                       `json:"app_id" example:"440"`
	Name         *string                       `json:"name" example:"Team Fortress 2"`
	StoreURLPath *string                       `json:"store_url_path" example:"app/440/Team_Fortress_2"`
	Type         string                        `json:"type" enum:"game,demo,mod" example:"game"`
	Links        *[]string                     `json:"links,omitempty,omitzero"`
	BasicInfo    *steamapi.StoreItem_BasicInfo `json:"basic_info,omitempty,omitzero"`
	Assets       *steamapi.StoreItem_Assets    `json:"assets,omitempty,omitzero"`
	Platforms    AppsPlatforms                 `json:"platforms,omitzero"`
	Release      AppsRelease                   `json:"release,omitzero"`
}

type AppsPlatforms struct {
	Windows      *bool `json:"windows" example:"true"`
	SteamOSLinux *bool `json:"steamos_linux" example:"true"`
	Mac          *bool `json:"mac" example:"false"`
	// TODO:
	// SteamVr
	//SteamDeckCOmpat
	//STeamOsCOmpat
}

type AppsRelease struct {
	SteamReleaseDate    time.Time `json:"steam_release_date,omitempty,omitzero"`
	OriginalReleaseDate time.Time `json:"original_release_date,omitempty,omitzero"`
}

type responseBody interface {
	searchSuggestionsResponse()
}

type raw steamapi.CStoreQuery_SearchSuggestions_Response

func (r *raw) searchSuggestionsResponse()          {}
func (r *AppsResponse) searchSuggestionsResponse() {}

type Response struct {
	Body responseBody
}
