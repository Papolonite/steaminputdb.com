package ping

import (
	"github.com/Alia5/steaminputdb.com/version"
	"github.com/go-fuego/fuego"
	"github.com/go-fuego/fuego/option"
)

type Ping struct {
	Service string `json:"service"`
	Version string `json:"version"`
}

func RegisterRoutes(s *fuego.Server) {
	fuego.Get(s, "/ping", Controller,
		option.Summary("Ping"),
		option.OverrideDescription(
			"Returns basic information about the service.",
		),
		option.Tags("Meta"),
	)
}

func Controller(_ fuego.ContextNoBody) (*Ping, error) {
	return &Ping{
		Service: "SteamInputDB.com",
		Version: version.Version,
	}, nil
}
