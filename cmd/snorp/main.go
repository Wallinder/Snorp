package main

import (
	"context"
	"snorp/internal/manager"
	"snorp/internal/state"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	session := state.NewState()

	manager.Controller(ctx, session)
}
