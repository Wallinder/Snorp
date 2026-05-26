package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"snorp/internal/receiver"
	"snorp/internal/server"
	"snorp/internal/state"
	"sync"
	"syscall"
)

func main() {
	session := state.NewState()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	var wg sync.WaitGroup

	httpServer := server.NewHttpServer(session)
	server.RunHttpServer(httpServer, &wg)

	session.ReadWebsocketErrors(ctx, &wg)
	session.Discord.StartWebsocket(ctx, &wg)

	receiver.StartDispatchReader(ctx, session, &wg)

	<-ctx.Done()
	server.Shutdown(ctx, httpServer)

	wg.Wait()
	slog.Info("snorp shutting down")
}
