package main

import (
	"log"
	"menial/config"
	event "menial/internal/event"
	"menial/internal/state"
	wss "menial/internal/websocket"

	"golang.org/x/net/websocket"
)

type Bot struct {
	StaticConfig config.StaticConfig
	Connection   *websocket.Conn
	State      state.SessionState
	Messages     chan []byte
}

func (b *Bot) Run() {
	b.State.Running = true

	b.Connection = wss.Connect(b.State.Metadata.Url)
	b.Messages = make(chan []byte)

	log.Println("Starting bot..")
	defer close(b.Messages)
	defer b.Connection.Close()

	go wss.Listen(b.Connection, b.Messages)

	event.MessageHandler(
		b.Connection,
		b.Messages,
		b.StaticConfig,
		&b.State,
	)
}

func main() {
	conf := config.Settings()
	var state state.SessionState
	state.UpdateMetadata(conf.Bot.Token, conf.Url.Gateway)
	bot := Bot{
		StaticConfig: conf,
		State:      state,
	}
	bot.Run()
}
