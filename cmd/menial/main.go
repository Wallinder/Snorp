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
	s.UpdateMetadata(s.Config.Bot.Token, s.Config.Bot.Gateway)

	topCtx := context.Background()
	defer close(s.Messages)

	websocketUrl := s.Metadata.Url

	for {
		ctx, cancel := context.WithCancel(topCtx)

		if s.Resume {
			websocketUrl = s.ReadyData.ResumeGatewayURL
		}

		conn, err := socket.Connect(ctx, websocketUrl)
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
