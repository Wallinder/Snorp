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

func (d *Discord) Start(ctx context.Context, wg *sync.WaitGroup) {
	wg.Go(func() {
		d.startWebsocket(ctx)
	})
}

func (d *Discord) startWebsocket(ctx context.Context) {
	for {
		if ctx.Err() != nil {
			return
		}
		if d.Websocket.ReconnectAttempts >= d.Websocket.MaxRetries {
			panic(ErrBackoffExceeded)
		}
		if time.Since(d.Websocket.LastAttempt) > d.Websocket.ResetAfter {
			d.Websocket.ReconnectAttempts = 0
		}
		d.Websocket.LastAttempt = time.Now()

		err := eventHandler(ctx, d)
		sendErr(d.ErrorChan, err)

		TotalDisconnects.Inc()
		d.Websocket.ReconnectAttempts++
	}
}

func sendErr(errChan chan error, err error) {
	select {
	case errChan <- err:
	default:
	}
}
