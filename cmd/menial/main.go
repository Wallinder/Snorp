package main

import (
	"context"
	"menial/config"
	"menial/internal/event"
	"menial/internal/socket"
	"menial/internal/state"
	"time"
)

func Run(s *state.SessionState) {
	defer close(s.Messages)

	websocketUrl := s.Metadata.Url

	for {
		ctx, cancel := context.WithCancel(context.Background())
		if s.Resume {
			websocketUrl = s.ReadyData.ResumeGatewayURL
		}

		conn, err := socket.Connect(ctx, websocketUrl)
		if err != nil {
			time.Sleep(60 * time.Second)
			continue
		}
		go socket.Listen(ctx, conn, s)

		event.MessageHandler(ctx, conn, s)
		cancel()

		time.Sleep(3 * time.Second)
	}
}

func main() {
	session := &state.SessionState{
		Config:   config.Settings(),
		Resume:   false,
		Messages: make(chan []byte),
	}
	session.UpdateMetadata(
		session.Config.Bot.Token,
		session.Config.Bot.Gateway,
	)
	Run(session)
}
