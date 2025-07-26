package main

import (
	"context"
	"log"
	"menial/config"
	"menial/internal/event"
	"menial/internal/state"
	socket "menial/internal/websocket"

	"github.com/coder/websocket"
)

type Bot struct {
	StaticConfig config.StaticConfig
	Connection   *websocket.Conn
	SessionState state.SessionState
	Messages     chan []byte
	Context      context.Context
}

func (b *Bot) Run() {
	b.Connection = socket.Connect(b.Context, b.SessionState.Metadata.Url)
	defer b.Connection.Close(1006, "Abornmal Closure")

	b.Messages = make(chan []byte)
	defer close(b.Messages)

	log.Println("Starting bot..")

	go socket.Listen(b.Context, b.Connection, b.Messages, &b.SessionState)

	for {
		err := event.MessageHandler(
			b.Context,
			b.Connection,
			b.Messages,
			b.StaticConfig,
			&b.SessionState,
		)
		if err != nil {
			b.Connection = socket.Connect(b.Context, b.SessionState.Metadata.Url)
			continue
		}
	}
}

func main() {
	conf := config.Settings()
	var sessionState state.SessionState
	sessionState.UpdateMetadata(conf.Bot.Token, conf.Url.Gateway)
	bot := Bot{
		StaticConfig: conf,
		SessionState: sessionState,
		Context:      context.Background(),
	}
	bot.Run()
}
