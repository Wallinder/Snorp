package program

import (
	"context"
	"log/slog"
	"sync"
)

type ServiceMon struct {
	Origin  string
	ErrChan chan error
}

func (app *Application) startErrorHandler(ctx context.Context, wg *sync.WaitGroup) {
	wg.Go(func() {
		for {
			select {

			case <-ctx.Done():
				return

			case err := <-app.Discord.Websocket.ErrorChan:
				slog.Error(err.Error(), "origin", "discord/ws")

			case err := <-app.ErrorChan:
				slog.Error(err.Error(), "origin", "internal")
			}
		}
	})
}
