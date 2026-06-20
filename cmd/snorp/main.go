package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"snorp/internal/program"
	"sync"
	"syscall"
)

func main() {
	app := program.NewApplication()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	var wg sync.WaitGroup

	app.InitDependencies(ctx)
	app.Start(ctx, &wg)

	<-ctx.Done()
	app.Stop(ctx, &wg)

	slog.Info("snorp stopped gracefully")
}
