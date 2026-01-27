package api

import "net/http"

type StatusWriter struct {
	http.ResponseWriter
	Status int
}

func (w *StatusWriter) WriteHeader(code int) {
	w.Status = code
	w.ResponseWriter.WriteHeader(code)
}
