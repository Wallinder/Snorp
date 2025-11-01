package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"snorp/config"
	"snorp/internal/event"
	"snorp/internal/metrics"
	"snorp/internal/sql"
	"snorp/internal/state"
	"time"
)

func Start(session *state.SessionState) {
	session.StartTime = time.Now()
	ctx := context.Background()

	metricAddr := fmt.Sprintf(":%d", session.MetricPort)
	session.MetricServer = &http.Server{
		Addr:           metricAddr,
		Handler:        metrics.MetricHandler(session),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go session.MetricServer.ListenAndServe()

	if session.Config.Postgresql.Enabled {
		session.Pool = sql.CreatePool(
			ctx,
			session.Config.Postgresql.ConnectionString,
		)
		defer session.Pool.Close()

		err := sql.InitDatabase(ctx, session.Pool)
		if err != nil {
			log.Fatalf("Error initializing db: %v", err)
		}
	}
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
	}
	session.InitHttpClient()
	session.UpdateMetadata()

	Start(session)
}
