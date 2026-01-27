package api

import (
	"context"
	"errors"
	"net/http"
	"os"

	"github.com/go-fuego/fuego"
)

func ErrorHandler(ctx context.Context, err error) error {

	var errWithStatus fuego.ErrorWithStatus
	if errors.As(err, &errWithStatus) {
		err = fuego.HandleHTTPError(ctx, err)
	}

	isDebug := os.Getenv("DEBUG") == "true" || os.Getenv("ENV") == "dev"

	instance := ""
	if uri, ok := ctx.Value("requestURI").(string); ok {
		instance = uri
	}
	var httpErr fuego.HTTPError
	if errors.As(err, &httpErr) {
		if httpErr.Status == 0 {
			httpErr.Status = httpErr.StatusCode()
		}
		if httpErr.Type == "" {
			httpErr.Type = "about:blank"
		}
		if httpErr.Title == "" {
			httpErr.Title = http.StatusText(httpErr.StatusCode())
			if httpErr.Title == "" {
				httpErr.Title = "HTTP Error"
			}
		}
		if httpErr.Instance == "" {
			httpErr.Instance = instance
		}
		if httpErr.Err == nil {
			httpErr.Err = err
		}

		if isDebug && httpErr.Detail == "" {
			httpErr.Detail = err.Error()
		}

		return httpErr
	}

	detail := "An internal error occurred"
	if isDebug {
		detail = err.Error()
	}

	return fuego.InternalServerError{
		Type:     "about:blank",
		Status:   http.StatusInternalServerError,
		Title:    "Internal Server Error",
		Detail:   detail,
		Instance: instance,
		Err:      err,
	}
}
