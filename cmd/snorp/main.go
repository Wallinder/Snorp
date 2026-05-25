package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"snorp/internal/server"
	"snorp/internal/state"
	"sync"
)

func main() {
	session := state.NewState()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	var wg sync.WaitGroup

	httpServer := server.NewHttpServer(session)
	server.RunHttpServer(&wg, httpServer)

	go session.ReadWebsocketErrors(ctx)
	session.Discord.StartWebsocket(ctx, &wg)

	<-ctx.Done()
	server.Shutdown(ctx, httpServer)

	wg.Wait()
	slog.Info("snorp shutting down")
}
