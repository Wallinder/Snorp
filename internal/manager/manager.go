package manager

import (
	"context"
	"snorp/internal/state"
	"time"
)

type Controller interface {
	start(context.Context)
}

func StartControllers(ctx context.Context, session *state.SessionState) {
	controllers := []Controller{
		&WebsocketController{
			Session:    session,
			MaxRetries: 3,
			ResetAfter: 30 * time.Second,
		},
	}
	for _, controller := range controllers {
		controller.start(ctx)
	}
}
