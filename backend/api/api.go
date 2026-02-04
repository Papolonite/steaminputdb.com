package api

import (
	"github.com/Alia5/steaminputdb.com/api/ping"
	"github.com/Alia5/steaminputdb.com/api/search"
	"github.com/Alia5/steaminputdb.com/api/steam"
	"github.com/danielgtaylor/huma/v2"
)

func RegisterAPI(a huma.API) {
	ping.RegisterRoutes(a)
	steam.RegisterRoutes(a)
	search.RegisterRoutes(a)
}
