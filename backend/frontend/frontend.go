package frontend

import (
	"embed"
	"io/fs"
	"net/http"
	"path"
)

//go:embed all:dist
var frontend embed.FS

func Handler(w http.ResponseWriter, r *http.Request) {
	filePath := path.Join("dist", r.URL.Path)
	if _, err := fs.Stat(frontend, filePath); err != nil {
		filePath = "dist"
	}
	http.ServeFileFS(w, r, frontend, filePath)
}
