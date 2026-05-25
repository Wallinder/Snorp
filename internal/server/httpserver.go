package server

import (
	"context"
	"net/http"
	"snorp/internal/state"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func RunHttpServer(wg *sync.WaitGroup, server *http.Server) {
	wg.Go(func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			state.LogAndExit("http panic", err, 1)
		}
	})
}

func Shutdown(ctx context.Context, server *http.Server) {
	newCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	err := server.Shutdown(newCtx)
	if err != nil {
		state.LogAndExit("failed to gracefully stop server", err, 1)
	}
}

func NewHttpServer(session *state.SessionState) *http.Server {
	return &http.Server{
		Addr:              ":8080",
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       5 * time.Second,
		MaxHeaderBytes:    1 << 20,
		Handler:           requestHandler(session),
	}
}

func requestHandler(session *state.SessionState) http.Handler {
	router := http.NewServeMux()

	router.Handle("GET /metrics", promhttp.Handler())

	router.HandleFunc("GET /readyz", func(w http.ResponseWriter, r *http.Request) {
		if !session.IsReady() {
			http.Error(w, "instance is not ready yet", http.StatusInternalServerError)
			return
		}
		w.Write([]byte("ready"))
	})

	return defaults(router)
}
