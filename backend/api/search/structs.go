package search

import (
	"github.com/Alia5/steaminputdb.com/api/search/configs"
	"github.com/Alia5/steaminputdb.com/api/search/games"
)

type Request struct {
	Body SearchRequestBody
}

type Response struct {
	Body SearchResponseBody
}

type SearchRequestBody struct {
	Limit        uint32 `json:"limit,omitempty,omitzero" max:"20" default:"10"`
	LimitGames   uint32 `json:"limit_games,omitempty,omitzero" default:"3" max:"5"`
	LimitConfigs uint32 `json:"limit_configs,omitempty,omitzero" default:"5" max:"15"`

	SearchTerm string `json:"search_term" required:"true" maxLength:"100"`
}

type SearchResponseBody struct {
	Games   []games.AppItem       `json:"games,omitempty"`
	Configs []configs.ConfigItem `json:"configs,omitempty"`
}
