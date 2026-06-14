package server

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Start(server *http.Server, wg *sync.WaitGroup) {
	wg.Go(func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	})
}

func Stop(ctx context.Context, server *http.Server) {
	newCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := server.Shutdown(newCtx)
	if err != nil {
		panic(err)
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

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Handle("GET /metrics", promhttp.Handler())

	return defaults(router)
}
