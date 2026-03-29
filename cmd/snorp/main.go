package main

import (
	"context"
	"snorp/internal/event"
	"snorp/internal/metrics"
	"snorp/internal/state"
)

func main() {
	session := state.NewState()

	go metrics.Collector(session.Metrics)

	ctx := context.Background()
	event.EventListener(ctx, session)
}
