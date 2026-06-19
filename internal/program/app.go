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
	Discord   *discord.Discord
	ErrorChan chan Errors
}

type Errors struct {
	Err    error
	Origin string
	Fatal  bool
}

func NewApplication() *Application {
	config, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	return &Application{
		Config:    config,
		Server:    server.NewHttpServer(),
		Client:    client.NewHttpClient(),
		ErrorChan: make(chan Errors),
		StartTime: time.Now(),
	}
}

func (app *Application) Start(ctx context.Context, wg *sync.WaitGroup) {
	server.Start(app.Server, wg)

	var err error
	app.Discord, err = discord.NewDiscord(
		app.Client,
		app.Config.Bot.Identity,
		app.Config.Bot.Api,
		app.Config.Bot.ApiVersion,
	)
	if err != nil {
		panic(err)
	}
	app.ErrorHandler(ctx, wg)
	app.Discord.StartWebsocket(ctx, wg)
}

func (app *Application) Stop(ctx context.Context, wg *sync.WaitGroup) {
	server.Stop(ctx, app.Server)
	wg.Wait()
}
