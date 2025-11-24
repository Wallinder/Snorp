package main

import (
	"context"
	"snorp/config"
	"snorp/internal/event"
	"snorp/internal/metrics"
	"snorp/internal/state"
	"time"
)

func Start(session *state.SessionState) {
	session.StartTime = time.Now()
	ctx := context.Background()

	session.DB = session.CreateConnection()

	session.UpdateMetadata()

	event.EventListener(ctx, session)
}

func main() {
	session := &state.SessionState{
		Config:     config.Settings(),
		Resume:     false,
		Messages:   make(chan []byte),
		MaxRetries: 3,
		MetricUri:  "/metrics",
		MetricPort: 8080,
		Jobs: state.Jobs{
			SteamNews:  make(map[string]bool),
			SteamSales: make(map[string]bool),
		},
	}
	go metrics.Collector(session)

	session.InitHttpClient()

	Start(session)
}
