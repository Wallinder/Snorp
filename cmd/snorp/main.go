package main

import (
	"context"
	"snorp/config"
	"snorp/internal/socket/event"
	"snorp/internal/state"
)

func Start(s *state.SessionState) {
	ctx := context.Background()
	event.EventListener(ctx, s)
}

func main() {
	session := &state.SessionState{
		Config:     config.Settings(),
		Resume:     false,
		Messages:   make(chan []byte),
		MaxRetries: 3,
	}
	session.InitHttpClient()
	state.UpdateMetadata(session)

	Start(session)
}
