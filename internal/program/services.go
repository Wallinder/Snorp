package program

import (
	"context"
	"log/slog"
	"snorp/internal/services/receiver"
	"sync"
)

type Services interface {
	Start(context.Context, *sync.WaitGroup)
	Name() string
}

func (app *Application) startServices(ctx context.Context, wg *sync.WaitGroup) {
	services := []Services{
		&receiver.DispatcherService{
			Discord: app.Discord,
			ErrChan: app.ErrorChan,
		},
		app.Discord,
	}

	for _, service := range services {
		slog.Info("started", "service", service.Name())
		service.Start(ctx, wg)
	}
}
