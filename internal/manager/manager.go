package manager

import (
	"context"
	"snorp/internal/state"
	"sync"
	"time"
)

type Controller interface {
	start(context.Context)
}

func StartControllers(ctx context.Context, wg *sync.WaitGroup, session *state.SessionState) {
	controllers := []Controller{
		&WebsocketController{
			Session:    session,
			MaxRetries: 3,
			ResetAfter: 30 * time.Second,
		},
	}
	for _, controller := range controllers {
		wg.Go(func() {
			go controller.start(ctx)
		})
	}
}
