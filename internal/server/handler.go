package server

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func RequestHandler() http.Handler {
	router := http.NewServeMux()

	router.Handle("/metrics", promhttp.Handler())

	return defaults(router)
}
