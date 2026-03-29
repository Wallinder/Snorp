package event

import (
	"context"
	"log/slog"
	"snorp/internal/state"
	"time"
)

func EventListener(ctx context.Context, session *state.SessionState) {
	const resetAfter = 30 * time.Second

	var attempts int
	var lastAttempt time.Time

	for {
		if attempts >= session.MaxRetries {
			slog.Error("backoff timer exceeded, exiting..")
			return
		}
		if time.Since(lastAttempt) > resetAfter {
			attempts = 0
		}
		lastAttempt = time.Now()

		newCtx, cancel := context.WithCancel(ctx)
		eventHandler(newCtx, cancel, session)

		session.Metrics.TotalDisconnects.Inc()
		attempts++
	}
}
