package main

import (
	"context"
	"net/http"
	"snorp/internal/event"
	"snorp/internal/state"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(":8080", nil); err != nil {
		state.LogAndExit("httpserver panic", err, 1)
	}

	ctx := context.Background()
	session := state.NewState()

	event.Listener(ctx, session)
}
