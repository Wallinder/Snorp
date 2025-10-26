package main

import (
	"context"
	"log"
	"snorp/config"
	"snorp/internal/event"
	"snorp/internal/sql"
	"snorp/internal/state"
)

func Start(session *state.SessionState) {
	ctx := context.Background()

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
	}
	session.InitHttpClient()
	session.UpdateMetadata()

	Start(session)
}
