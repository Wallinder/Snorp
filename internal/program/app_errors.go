package program

import (
	"context"
	"log/slog"
	"sync"
)

func (app *Application) errorHandler(ctx context.Context, wg *sync.WaitGroup) {
	wg.Go(func() {
		for {
			select {

			case <-ctx.Done():
				return

			case err := <-app.Discord.ErrorChan:
				slog.Error(err.Error(), "origin", "discord")

			case err := <-app.Services.Dispatcher.ErrChan:
				slog.Error(err.Error(), "origin", "dispatcher")
			}
		}
	})
}
