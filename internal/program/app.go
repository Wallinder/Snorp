package program

import (
	"context"
	"net/http"
	"snorp/config"
	"snorp/internal/client"
	"snorp/internal/server"
	"snorp/pkg/discord"
	"sync"
	"time"
)

type Application struct {
	StartTime time.Time
	Config    *config.Config
	Server    *http.Server
	Client    *http.Client
	Discord   *discord.DiscordService
	ErrorChan chan error
}

func NewApplication(ctx context.Context, wg *sync.WaitGroup) *Application {
	config, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	client := client.NewHttpClient()
	errChan := make(chan error)

	discord, err := discord.NewDiscord(
		client,
		config.Bot.Identity,
		config.Bot.Api,
		config.Bot.ApiVersion,
		errChan,
	)
	if err != nil {
		panic(err)
	}

	return &Application{
		Config:    config,
		Server:    server.NewHttpServer(),
		Client:    client,
		StartTime: time.Now(),
		Discord:   discord,
		ErrorChan: errChan,
	}
}

func (app *Application) Start(ctx context.Context, wg *sync.WaitGroup) {
	server.Start(app.Server, wg)
	app.startErrorHandler(ctx, wg)
	app.startServices(ctx, wg)
}

func (app *Application) Stop(ctx context.Context, wg *sync.WaitGroup) {
	server.Stop(ctx, app.Server)
	wg.Wait()
}
