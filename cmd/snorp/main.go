package main

import (
	"context"
	"snorp/config"
	"snorp/internal/socket/event"
	"snorp/internal/state"
)

func Start(session *state.SessionState) {
	ctx := context.Background()
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
