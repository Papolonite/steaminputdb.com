package routes

import (
	"github.com/go-fuego/fuego"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func RegisterMetrics(s *fuego.Server) {
	fuego.Handle(s, "/metrics", promhttp.Handler())
}
