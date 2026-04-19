package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"snorp/internal/manager"
	"snorp/internal/server"
	"snorp/internal/state"
	"sync"
)

func main() {
	session := state.NewState()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	var wg sync.WaitGroup

	httpServer := server.NewHttpServer()
	server.RunHttpServer(&wg, httpServer)

	manager.StartControllers(ctx, &wg, session)

	<-ctx.Done()
	server.Shutdown(ctx, httpServer)

	wg.Wait()
	slog.Info("snorp shutting down")
}
