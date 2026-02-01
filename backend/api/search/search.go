package search

import (
	"github.com/Alia5/steaminputdb.com/api/search/configs"
	"github.com/Alia5/steaminputdb.com/api/search/games"
	"github.com/danielgtaylor/huma/v2"
)

func RegisterRoutes(a huma.API) {
	configs.RegisterRoutes(a)
	games.RegisterRoutes(a)
}
