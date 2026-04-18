package manager

import (
	"context"
	"snorp/internal/state"
	"time"
)

type Controller interface {
	start(context.Context, *state.SessionState)
}

func StartControllers(ctx context.Context, session *state.SessionState) {
	controllers := []Controller{
		&WebsocketController{
			ResetAfter: 30 * time.Second,
		},
	}
	for _, controller := range controllers {
		controller.start(ctx, session)
	}
}
