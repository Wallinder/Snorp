package manager

import (
	"context"
	"log/slog"
	"snorp/internal/event"
	"snorp/internal/state"
	"time"
)

type WebsocketController struct {
	ResetAfter  time.Duration
	Attempts    int
	LastAttempt time.Time
}

func (wc *WebsocketController) start(ctx context.Context, session *state.SessionState) {
	for {
		if ctx.Err() != nil {
			return
		}
		if wc.Attempts >= session.MaxRetries {
			slog.Error("backoff timer exceeded, exiting..")
			return
		}
		if time.Since(wc.LastAttempt) > wc.ResetAfter {
			wc.Attempts = 0
		}
		wc.LastAttempt = time.Now()

		newCtx, cancel := context.WithCancel(ctx)
		event.EventHandler(newCtx, cancel, session)

		TotalDisconnects.Inc()
		wc.Attempts++
	}
}
