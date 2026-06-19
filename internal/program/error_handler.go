package program

import (
	"context"
	"log/slog"
	"os"
	"sync"
)

func (app *Application) ErrorHandler(ctx context.Context, wg *sync.WaitGroup) {
	wg.Go(func() {
		for {
			select {

			case <-ctx.Done():
				return

			case err := <-app.Discord.Websocket.ErrorChan:
				slog.Error(err.Error(), "origin", "discord/ws")

			case msg := <-app.ErrorChan:
				slog.Error(msg.Err.Error(), "origin", msg.Origin)

				if msg.Fatal {
					os.Exit(1)
				}
			}
		}
	})
}
