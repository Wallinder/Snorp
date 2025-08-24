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
	ctx, cancel := context.WithCancel(b.Context)

	b.Messages = make(chan []byte)
	defer close(b.Messages)

	b.StopChannel = make(chan bool)
	defer close(b.StopChannel)

	for {
		wss := b.SessionState.Metadata.Url

		b.Connection = socket.Connect(ctx, wss)
		log.Printf("Connected to socket: %s\n", wss)

		go socket.Listen(ctx, b.Connection, b.Messages, &b.SessionState)

		event.MessageHandler(ctx,
			b.Connection,
			b.Messages,
			b.StaticConfig,
			&b.SessionState,
		)
		cancel()
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
