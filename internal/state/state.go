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
	Discord   *discord.Discord
	ErrorChan chan SessionError
	StartTime time.Time
	Config    *config.Config
	Client    *http.Client
}

type SessionError struct {
	Err    error
	Origin string
	Fatal  bool
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
		state.ErrorChan <- SessionError{Origin: "discord", Fatal: true, Err: err}
	}
	return state
}

func newDefaultState() *SessionState {
	config, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	return &SessionState{
		Config:    config,
		Client:    client.NewHttpClient(),
		StartTime: time.Now(),
		ErrorChan: make(chan SessionError),
	}
}

func (s *SessionState) ErrorHandler(ctx context.Context, wg *sync.WaitGroup) {
	wg.Go(func() {
		for {
			select {

			case <-ctx.Done():
				return

			case err := <-s.Discord.Websocket.ErrorChan:
				slog.Error(err.Error(), "origin", "discord/ws")

			case msg := <-s.ErrorChan:
				slog.Error(msg.Err.Error(), "origin", msg.Origin)

				if msg.Fatal {
					os.Exit(1)
				}
			}
		}
	})
}
