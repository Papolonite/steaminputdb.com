package api

import (
	"github.com/Alia5/steaminputdb.com/api/login"
	"github.com/Alia5/steaminputdb.com/api/ping"
	"github.com/Alia5/steaminputdb.com/api/search"
	"github.com/danielgtaylor/huma/v2"
)

func RegisterAPI(a huma.API) {
	ping.RegisterRoutes(a)
	login.RegisterRoutes(a)
	search.RegisterRoutes(a)
}
