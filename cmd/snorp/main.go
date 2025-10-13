package main

import (
	"context"
	"snorp/config"
	"snorp/internal/event"
	"snorp/internal/sql"
	"snorp/internal/state"
)

func Start(session *state.SessionState) {
	ctx := context.Background()

	if session.Config.Postgresql.Enabled {
		session.ConnectionPool = sql.CreatePool(
			ctx,
			session.Config.Postgresql.ConnectionString,
		)
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
