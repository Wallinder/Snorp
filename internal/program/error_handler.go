package program

import (
	"context"
	"log/slog"
	"sync"
)

func (app *Application) startErrorHandler(ctx context.Context, wg *sync.WaitGroup) {
	wg.Go(func() {
		for {
			select {

			case <-ctx.Done():
				return

			case err := <-app.ErrorChan:
				slog.Error(err.Error(), "origin", "internal")
			}
		}
	})
}
