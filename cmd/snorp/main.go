package main

import (
	"context"
	"net/http"
	"snorp/internal/event"
	"snorp/internal/state"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(":8080", nil)

	session := state.NewState()

	event.Listener(ctx, session)
}
