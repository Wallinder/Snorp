package main

import (
	"context"
	"log"
	"menial/config"
	"menial/internal/event"
	"menial/internal/state"
	socket "menial/internal/websocket"
	"time"

	"github.com/coder/websocket"
)

type Bot struct {
	StaticConfig config.StaticConfig
	Connection   *websocket.Conn
	SessionState state.SessionState
	StopChannel  chan bool
	Messages     chan []byte
	Context      context.Context
}

func (b *Bot) Run() {
	b.Messages = make(chan []byte)
	defer close(b.Messages)

	b.StopChannel = make(chan bool)
	defer close(b.StopChannel)

	for {
		wss := b.SessionState.Metadata.Url

		b.Connection = socket.Connect(b.Context, wss)
		log.Printf("Connected to socket: %s\n", wss)

		go socket.Listen(b.Context, b.Connection, b.Messages, b.StopChannel, &b.SessionState)

		event.MessageHandler(
			b.Context,
			b.Connection,
			b.Messages,
			b.StaticConfig,
			&b.SessionState,
		)
		b.Connection.Close(1006, "Normal Closure")

		time.Sleep(10 * time.Second)
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
