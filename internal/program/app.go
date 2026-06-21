package program

import (
	"context"
	"net/http"
	"snorp/config"
	"snorp/internal/client"
	"snorp/internal/server"
	"snorp/internal/services/receiver"
	"snorp/pkg/discord"
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Application struct {
	StartTime   time.Time
	Config      *config.Config
	Server      *http.Server
	Client      *http.Client
	PgPool      *pgxpool.Pool
	Discord     *discord.Discord
	Services    Services
	StorageType string
	Storage     Storage
}

type Services struct {
	Dispatcher *receiver.DispatcherService
}

func NewApplication() *Application {
	config, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	return &Application{
		Config:      config,
		Server:      server.NewHttpServer(),
		Client:      client.NewHttpClient(),
		StartTime:   time.Now(),
		StorageType: "file",
	}
}

func (app *Application) InitDependencies(ctx context.Context) {
	var err error
	if app.Config.Postgres.Enabled {
		app.StorageType = "postgres"
	}

	app.Storage, err = app.NewStorage(ctx, app.StorageType)
	if err != nil {
		panic(err)
	}

	app.Discord, err = discord.NewDiscord(
		app.Client,
		app.Config.Bot.Identity,
		app.Config.Bot.Api,
		app.Config.Bot.ApiVersion,
	)
	if err != nil {
		panic(err)
	}
}

func (app *Application) Start(ctx context.Context, wg *sync.WaitGroup) {
	app.Services = Services{
		Dispatcher: receiver.NewDispatchService(app.Discord),
	}

	server.Start(app.Server, wg)
	app.errorHandler(ctx, wg)

	app.Services.Dispatcher.Start(ctx, wg)
	app.Discord.Start(ctx, wg)
}

func (app *Application) Stop(ctx context.Context, wg *sync.WaitGroup) {
	server.Stop(ctx, app.Server)
	if app.PgPool != nil {
		app.PgPool.Close()
	}
	wg.Wait()
}
