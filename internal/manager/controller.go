package manager

import (
	"context"
	"log/slog"
	"snorp/internal/event"
	"snorp/internal/server"
	"snorp/internal/state"
	"time"
)

func Controller(ctx context.Context, session *state.SessionState) {
	const resetAfter = 30 * time.Second

	var attempts int
	var lastAttempt time.Time

	httpServer := server.NewHttpServer()
	go server.RunHttpServer(httpServer)

	for {
		select {
		case <-ctx.Done():
			slog.Info("controller shutting down")
			err := httpServer.Shutdown(context.Background())
			if err != nil {
				state.LogAndExit("failed to gracefully stop server", err, 1)
			}
			return

		default:
			if attempts >= session.MaxRetries {
				slog.Error("backoff timer exceeded, exiting..")
				return
			}
			if time.Since(lastAttempt) > resetAfter {
				attempts = 0
			}
			lastAttempt = time.Now()

			newCtx, cancel := context.WithCancel(ctx)
			event.EventHandler(newCtx, cancel, session)

			TotalDisconnects.Inc()
			attempts++
		}
	}
}
