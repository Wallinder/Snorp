package discord

import (
	"context"
	"errors"
	"sync"
	"time"
)

var (
	ErrBackoffExceeded  = errors.New("backoff timer exceeded")
	ErrContextCancelled = errors.New("context cancelled")
)

func (d *Discord) StartWebsocket(ctx context.Context, wg *sync.WaitGroup) {
	wg.Go(func() {
		go d.start(ctx)
	})
}

func (d *Discord) start(ctx context.Context) {
	for {
		if ctx.Err() != nil {
			return
		}
		if d.Websocket.ReconnectAttempts > d.Websocket.MaxRetries {
			panic(ErrBackoffExceeded)
		}
		if time.Since(d.Websocket.LastAttempt) > d.Websocket.ResetAfter {
			d.Websocket.ReconnectAttempts = 0
		}
		d.Websocket.LastAttempt = time.Now()

		d.Websocket.ErrorChan <- eventHandler(ctx, d)

		TotalDisconnects.Inc()
		d.Websocket.ReconnectAttempts++
	}
}
