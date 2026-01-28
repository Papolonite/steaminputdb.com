package routes

import (
	"encoding/json"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func UnregisteredMiddleware(next http.Handler) http.Handler {
	mux, ok := next.(*http.ServeMux)
	if ok {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, pattern := mux.Handler(r)
			if pattern == "" {
				methods := []string{
					http.MethodGet,
					http.MethodPost,
					http.MethodPut,
					http.MethodDelete,
					http.MethodPatch,
					http.MethodHead,
					http.MethodOptions,
					http.MethodTrace,
				}
				pathExists := false
				for _, method := range methods {
					testReq := &http.Request{Method: method, URL: r.URL, Host: r.Host}
					if _, p := mux.Handler(testReq); p != "" {
						pathExists = true
						break
					}
				}
				if pathExists {
					hErr := huma.Error405MethodNotAllowed("Method not allowed")
					b, err := json.Marshal(hErr)
					if err != nil {
						http.Error(w, "Internal server error", http.StatusInternalServerError)
						return
					}
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(hErr.GetStatus())
					_, _ = w.Write(b)
				} else {
					hErr := huma.Error404NotFound("Resource not found")
					b, err := json.Marshal(hErr)
					if err != nil {
						http.Error(w, "Internal server error", http.StatusInternalServerError)
						return
					}
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(hErr.GetStatus())
					_, _ = w.Write(b)
				}
				return
			}
			mux.ServeHTTP(w, r)
		})
	}
	return next
}
