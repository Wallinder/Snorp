package main

import (
	"context"
	"menial/config"
	"menial/internal/socket/event"
	"menial/internal/state"
)

func Run(s *state.SessionState) {
	ctx := context.Background()
	event.EventListener(ctx, s)
}

func main() {
	session := &state.SessionState{
		Config:     config.Settings(),
		Resume:     false,
		Messages:   make(chan []byte),
		MaxRetries: 5,
	}
	session.UpdateMetadata(
		session.Config.Bot.Token,
		session.Config.Bot.Gateway,
	)
	Run(session)
}
