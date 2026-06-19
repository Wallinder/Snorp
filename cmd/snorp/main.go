package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"snorp/internal/program"
	"snorp/internal/receiver"
	"sync"
	"syscall"
)

func main() {
	app := program.NewApplication()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	var wg sync.WaitGroup

	app.Start(ctx, &wg)
	receiver.StartDispatchReader(ctx, app, &wg)

	<-ctx.Done()
	app.Stop(ctx, &wg)

	slog.Info("snorp stopped gracefully")
}
