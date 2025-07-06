package main

import (
	"log"
	"menial/internal/config"
	event "menial/internal/events"
	message "menial/internal/messages"

	"github.com/gorilla/websocket"
)

type Bot struct {
	Configuration  config.Setting
	Metadata       event.Metadata
	Websocket      *websocket.Conn
	MessageChannel chan map[string]any
	Running        bool
}

func (b *Bot) Run() {
	b.Running = true
	log.Println("Starting bot..")
	defer close(b.MessageChannel)
	defer b.Websocket.Close()
	go message.Receiver(
		b.Websocket,
		b.MessageChannel,
	)
	message.Handler(
		b.Websocket,
		b.MessageChannel,
		b.Configuration,
	)
}

func main() {
	conf := config.Settings()
	metadata := event.GetGateway(conf.Env.Gateway, conf.Env.Token)
	bot := Bot{
		Configuration:  conf,
		Metadata:       metadata,
		Websocket:      event.Connect(metadata.Url),
		MessageChannel: make(chan map[string]any),
	}
	bot.Run()
}
