package event

import (
	"context"
	"log"
	"snorp/internal/state"
	"time"
)

const APIversion = "10"

func EventListener(ctx context.Context, session *state.SessionState) {
	const resetAfter = 30 * time.Second

	var attempts int
	var lastAttempt time.Time

	for {
		if attempts == 3 {
			session.Resume = false
		}
		if attempts >= session.MaxRetries {
			log.Fatal("Backoff timer exceeded, exiting..")
			return
		}
		if time.Since(lastAttempt) > resetAfter {
			attempts = 0
		}
		lastAttempt = time.Now()

		newCtx, cancel := context.WithCancel(ctx)
		EventHandler(newCtx, cancel, session)

		attempts++
	}
}
