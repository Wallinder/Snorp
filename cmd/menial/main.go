package main

import (
	"context"
	"log"
	"menial/config"
	"menial/internal/event"
	"menial/internal/state"
	socket "menial/internal/websocket"
	"time"
)

type Bot struct {
	StaticConfig config.StaticConfig
	SessionState state.SessionState
	Messages     chan []byte
}

func (b *Bot) Run() {
	topContext := context.Background()

	b.Messages = make(chan []byte)
	defer close(b.Messages)

	for {
		ctx, cancel := context.WithCancel(topContext)

		conn, err := socket.Connect(ctx, b.SessionState.Metadata.Url)
		if err != nil {
			log.Println(err)
			time.Sleep(60 * time.Second)
			continue
		}

		go socket.Listen(ctx, conn, b.Messages, &b.SessionState)

		if b.SessionState.Resume {
			event.ResumeConnection(ctx, conn, b.StaticConfig.Bot.Token, &b.SessionState)
		}

		event.MessageHandler(ctx, conn, b.Messages, b.StaticConfig, &b.SessionState)
		cancel()

		conn.Close(1006, "Normal Closure")
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
	}
	bot.Run()
}
