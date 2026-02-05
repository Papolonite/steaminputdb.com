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
	Limit        uint32 `json:"limit" max:"20"`
	LimitGames   uint32 `json:"limit_games" default:"3" max:"5"`
	LimitConfigs uint32 `json:"limit_configs" default:"5" max:"15"`

	SearchTerm string `json:"search_term" required:"true" maxLength:"100"`
}

type SearchResponseBody struct {
	Games   []games.AppsSearchItem       `json:"games,omitempty"`
	Configs []configs.ConfigResponseItem `json:"configs,omitempty"`
}
