package state

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"snorp/config"
	"snorp/internal/client"
	"snorp/pkg/discord"
	"sync"
	"time"
)

type SessionState struct {
	Discord     *discord.Discord
	StartTime   time.Time
	Config      *config.Config
	Client      *http.Client
	CommandsDir string
	Commands    []discord.ApplicationCommand
}

func NewState() *SessionState {
	state := newDefaultState()

	var err error
	state.Discord, err = discord.NewDiscord(
		state.Client,
		state.Config.Bot.Identity,
		state.Config.Bot.Api,
		state.Config.Bot.ApiVersion,
	)
	if err != nil {
		LogAndExit("unable to initialize discord", "discord", err)
	}
	//state.setCommands()

	return state
}

func newDefaultState() *SessionState {
	config, err := config.NewConfig()
	if err != nil {
		LogAndExit("unable to load configuration", "config", err)
	}
	return &SessionState{
		Config:      config,
		Client:      client.NewHttpClient(),
		CommandsDir: "./commands",
		StartTime:   time.Now(),
	}
}

func LogAndExit(msg string, component string, err error) {
	slog.Error(msg, component, err)
	os.Exit(1)
}

func (s *SessionState) ErrorHandler(ctx context.Context, wg *sync.WaitGroup) {
	wg.Go(func() {
		for {
			select {
			case <-ctx.Done():
				return
			case err := <-s.Discord.Websocket.ErrorChan:
				slog.Error("discord", "websocket", err)
			}
		}
	})
}

func (s *SessionState) Start(ctx context.Context, wg *sync.WaitGroup) {
	s.ErrorHandler(ctx, wg)
	s.Discord.StartWebsocket(ctx, wg)
}
