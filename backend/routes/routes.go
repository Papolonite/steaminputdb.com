package routes

import "github.com/go-fuego/fuego"

func Register(f *fuego.Server, a *fuego.Server, m *fuego.Server) {
	RegisterFrontend(f)
	RegisterAPI(a)
	RegisterMetrics(m)
}
