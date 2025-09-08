package main

import (
	"context"
	"log"
	"menial/config"
	"menial/internal/socket/event"
	"menial/internal/state"
	"time"
)

func Run(s *state.SessionState) {

	ctx := context.Background()

	const resetAfter = 30 * time.Second

	var attempts int
	var lastAttempt time.Time

	for {
		if attempts >= s.MaxRetries {
			log.Fatal("Backoff timer exceeded, exiting..")
			return
		}
		if time.Since(lastAttempt) > resetAfter {
			attempts = 0
		}
		lastAttempt = time.Now()

		newCtx, cancel := context.WithCancel(ctx)
		event.EventListener(newCtx, cancel, s)

		attempts++
	}
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
