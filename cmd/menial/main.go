package main

import (
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
}

func (b *Bot) Run() {
	b.Connection = socket.Connect(b.SessionState.Metadata.Url)
	defer b.Connection.Close(1006, "Abornmal Closure")

	b.Messages = make(chan []byte)
	defer close(b.Messages)

	log.Println("Starting bot..")

	go socket.Listen(b.Connection, b.Messages)

	for {
		event.MessageHandler(
			b.Connection,
			b.Messages,
			b.StaticConfig,
			&b.SessionState,
		)
	}
}

func main() {
	conf := config.Settings()
	var sessionState state.SessionState
	sessionState.UpdateMetadata(conf.Bot.Token, conf.Url.Gateway)
	bot := Bot{
		StaticConfig: conf,
		SessionState: sessionState,
	}
	bot.Run()
}
