package manager

import (
	"context"
	"log/slog"
	"snorp/internal/event"
	"snorp/internal/state"
	"time"
)

type WebsocketController struct {
	Session           *state.SessionState
	MaxRetries        int
	ResetAfter        time.Duration
	ReconnectAttempts int
	LastAttempt       time.Time
}

func (wc *WebsocketController) start(ctx context.Context) {
	for {
		if ctx.Err() != nil {
			return
		}
		if wc.ReconnectAttempts >= wc.MaxRetries {
			slog.Error("backoff timer exceeded, exiting..")
			return
		}
		if time.Since(wc.LastAttempt) > wc.ResetAfter {
			wc.ReconnectAttempts = 0
		}
		wc.LastAttempt = time.Now()

		event.EventHandler(ctx, wc.Session)

		TotalDisconnects.Inc()
		wc.ReconnectAttempts++
	}
}
