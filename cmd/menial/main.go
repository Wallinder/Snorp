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
	Session      state.Session
	Messages     chan []byte
}

func (b *Bot) Run() {
	b.Session.Running = true

	b.Connection = wss.Connect(b.Session.Metadata.Url)
	b.Messages = make(chan []byte)

	log.Println("Starting bot..")
	defer close(b.Messages)
	defer b.Connection.Close()

	go wss.Listen(b.Connection, b.Messages)

	event.MessageHandler(
		b.Connection,
		b.Messages,
		b.StaticConfig,
		&b.Session,
	)
}

func main() {
	conf := config.Settings()
	var session state.Session
	session.UpdateMetadata(conf.Bot.Token, conf.Url.Gateway)
	bot := Bot{
		StaticConfig: conf,
		Session:      session,
	}
	bot.Run()
}
