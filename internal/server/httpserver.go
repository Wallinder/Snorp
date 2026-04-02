package server

import (
	"net/http"
	"snorp/internal/state"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func RunHttpServer(server *http.Server) {
	if err := server.ListenAndServe(); err != nil {
		state.LogAndExit("http panic", err, 1)
	}
}

func NewHttpServer() *http.Server {
	return &http.Server{
		Addr:              ":8080",
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       5 * time.Second,
		MaxHeaderBytes:    1 << 20,
		Handler:           requestHandler(),
	}
}

func requestHandler() http.Handler {
	router := http.NewServeMux()

	router.Handle("/metrics", promhttp.Handler())

	return defaults(router)
}
