package api

import (
	"net/http"

	"github.com/go-fuego/fuego"
	"github.com/go-fuego/fuego/option"
)

func RegisterCatchAll(s *fuego.Server) {
	fuego.Get(s, "/", func(_ fuego.ContextNoBody) (any, error) {
		return nil, fuego.NotFoundError{
			Detail: "The requested path was not found",
		}
	}, option.Hide())
	fuego.All(s, "/", func(_ fuego.ContextNoBody) (any, error) {
		return nil, fuego.HTTPError{
			Status: http.StatusMethodNotAllowed,
			Detail: "The requested method is not allowed",
		}
	}, option.Hide())
}
