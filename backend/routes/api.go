package routes

import (
	"github.com/Alia5/steaminputdb.com/api"
	"github.com/Alia5/steaminputdb.com/api/ping"
	"github.com/go-fuego/fuego"
)

func RegisterAPI(s *fuego.Server) {
	api.RegisterCatchAll(s)
	ping.RegisterRoutes(s)
}
