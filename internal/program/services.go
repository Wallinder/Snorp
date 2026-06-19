package program

import (
	"context"
	"snorp/internal/services/receiver"
	"sync"
)

type Services interface {
	Start(context.Context, *sync.WaitGroup)
	ServiceName() string
}

func (app *Application) startServices(ctx context.Context, wg *sync.WaitGroup) {
	services := []Services{
		&receiver.DispatcherService{
			Discord: app.Discord,
			ErrChan: app.ErrorChan,
		},
	}

	for _, service := range services {
		service.Start(ctx, wg)
	}
}
