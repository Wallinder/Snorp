package main

import (
	"log"
	"menial/internal/config"
	event "menial/internal/events"
	wss "menial/internal/websocket"

	"golang.org/x/net/websocket"
)

type Bot struct {
	Configuration config.Setting
	Connection    *websocket.Conn
	Messages      chan []byte
	EventStruct   event.EventStructs
	Running       bool
}

func (b *Bot) Run() {
	b.Running = true
	b.Connection = wss.Connect(b.Configuration.Metadata.Url)
	b.Messages = make(chan []byte)

	log.Println("Starting bot..")
	defer close(b.Messages)
	defer b.Connection.Close()

	go wss.Listen(b.Connection, b.Messages)

	go event.MessageHandler(
		b.Connection,
		b.Messages,
		&b.EventStruct,
	)
}

func main() {
	conf := config.Settings()
	bot := Bot{
		Configuration: conf,
	}
	bot.Run()
}
