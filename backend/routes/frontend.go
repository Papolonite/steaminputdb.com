package routes

import (
	"github.com/Alia5/steaminputdb.com/frontend"
	"github.com/go-fuego/fuego"
)

func RegisterFrontend(s *fuego.Server) {
	fuego.GetStd(s, "/", frontend.Handler)
}
