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
	Messages     chan []byte
	TopContext   context.Context
}

func (b *Bot) Run() {
	b.Messages = make(chan []byte)
	defer close(b.Messages)

	for {
		ctx, cancel := context.WithCancel(b.TopContext)

		wss := b.SessionState.Metadata.Url
		b.Connection = socket.Connect(ctx, wss)

		log.Printf("Connected to socket: %s\n", wss)

		if b.SessionState.Resume {
			event.ResumeConnection(ctx, b.Connection, b.StaticConfig.Bot.Token, &b.SessionState)
		}

		go socket.Listen(ctx, b.Connection, b.Messages, &b.SessionState)
		event.MessageHandler(ctx, b.Connection, b.Messages, b.StaticConfig, &b.SessionState)

		b.Connection.Close(1006, "Normal Closure")
		cancel()

		time.Sleep(30 * time.Second)
	}
}

func main() {
	conf := config.Settings()
	sessionState := state.SessionState{
		Resume: false,
	}
	sessionState.UpdateMetadata(conf.Bot.Token, conf.Bot.Gateway)
	bot := Bot{
		StaticConfig: conf,
		SessionState: sessionState,
		TopContext:   context.Background(),
	}
	bot.Run()
}
