package main

import (
	"context"
	"menial/config"
	"menial/internal/event"
	"menial/internal/state"
	socket "menial/internal/websocket"
	"time"
)

func Run(s *state.SessionState) {
	s.UpdateMetadata(s.Config.Bot.Token, s.Config.Bot.Gateway)

	topContext := context.Background()
	defer close(s.Messages)

	for {
		ctx, cancel := context.WithCancel(topContext)

		conn, err := socket.Connect(ctx, s.Metadata.Url)
		if err != nil {
			time.Sleep(60 * time.Second)
			continue
		}

		go socket.Listen(ctx, conn, s.Messages)

		event.MessageHandler(ctx, conn, s)
		cancel()

		time.Sleep(3 * time.Second)
	}
}

func main() {
	conf := config.Settings()
	session := &state.SessionState{
		Config:   conf,
		Resume:   false,
		Messages: make(chan []byte),
	}
	Run(session)
}
