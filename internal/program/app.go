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
)

type Application struct {
	StartTime time.Time
	Config    *config.Config
	Server    *http.Server
	Client    *http.Client
	Discord   *discord.Discord
	Services  Services
}

type Services struct {
	Dispatcher *receiver.DispatcherService
}

func NewApplication() *Application {
	config, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	client := client.NewHttpClient()

	discord, err := discord.NewDiscord(
		client,
		config.Bot.Identity,
		config.Bot.Api,
		config.Bot.ApiVersion,
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
		Services: Services{
			Dispatcher: receiver.NewDispatchService(discord),
		},
	}
}

func (app *Application) Start(ctx context.Context, wg *sync.WaitGroup) {
	server.Start(app.Server, wg)
	app.errorHandler(ctx, wg)

	app.Services.Dispatcher.Start(ctx, wg)
	app.Discord.Start(ctx, wg)
}

func (app *Application) Stop(ctx context.Context, wg *sync.WaitGroup) {
	server.Stop(ctx, app.Server)
	wg.Wait()
}
