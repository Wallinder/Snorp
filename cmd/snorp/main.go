package main

import (
	"context"
	"fmt"
	"net/http"
	"snorp/config"
	"snorp/internal/event"
	"snorp/internal/metrics"
	"snorp/internal/state"
	"time"

	"gorm.io/gorm"
)

func Start(session *state.SessionState) {
	session.StartTime = time.Now()
	ctx := context.Background()

	session.DB = state.CreateConnection(
		session.Config.Postgresql.ConnectionString,
		&session.DBSettings,
	)
	session.InitDatabase()

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
	}
	session.MetricServer = &http.Server{
		Addr:           fmt.Sprintf(":%d", session.MetricPort),
		Handler:        metrics.MetricHandler(session),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go session.MetricServer.ListenAndServe()

	session.DBSettings = state.DBSettings{
		GormConfig:      &gorm.Config{},
		MaxIdleConns:    10,
		MaxOpenConns:    100,
		ConnMaxLifetime: time.Hour,
	}
	session.InitHttpClient()

	Start(session)
}
