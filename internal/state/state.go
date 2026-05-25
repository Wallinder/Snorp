package state

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"snorp/pkg/discord"
	"time"
)

type SessionState struct {
	Discord     *discord.Discord
	StartTime   time.Time
	Config      *Config
	Client      *http.Client
	CommandsDir string
	Commands    []discord.ApplicationCommand
}

func NewState() *SessionState {
	state := newDefaultState()

	var err error
	state.Discord, err = discord.NewDiscord(
		state.Client,
		&state.Config.Bot.Identity,
		state.Config.Bot.Api,
		state.Config.Bot.ApiVersion,
	)
	if err != nil {
		LogAndExit("unable to initialize discord", err, 1)
	}
	//state.setCommands()

	return state
}

func newDefaultState() *SessionState {
	return &SessionState{
		Config:      NewConfig(),
		Client:      newHttpClient(),
		CommandsDir: "./commands",
		StartTime:   time.Now(),
	}
}

func newHttpClient() *http.Client {
	return &http.Client{
		CheckRedirect: nil,
		Timeout:       5 * time.Second,
	}
}

func LogAndExit(msg string, err error, exitcode int) {
	slog.Error(msg, "error", err)
	os.Exit(exitcode)
}

func (s *SessionState) ReadWebsocketErrors(ctx context.Context) {
	select {
	case <-ctx.Done():
		return
	case err := <-s.Discord.Websocket.ErrorChan:
		slog.Error("websocket", "error", err)
	}
}

func (s *SessionState) IsReady() bool {
	if s.Discord.ReadyData != nil {
		return true
	}
	return false
}
