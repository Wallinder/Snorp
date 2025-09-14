package main

import (
	"context"
	"net/http"
	"snorp/config"
	"snorp/internal/socket/event"
	"snorp/internal/state"
	"time"
)

func Run(s *state.SessionState) {
	ctx := context.Background()
	event.EventListener(ctx, s)
}

func main() {
	session := &state.SessionState{
		Config:     config.Settings(),
		Resume:     false,
		Messages:   make(chan []byte),
		MaxRetries: 5,
		Client: &http.Client{
			Transport: &http.Transport{
				MaxIdleConns:       10,
				IdleConnTimeout:    10 * time.Second,
				DisableCompression: true,
			},
			CheckRedirect: nil,
			Timeout:       time.Duration(10 * time.Second),
		},
	}
	session.UpdateMetadata(
		session.Config.Bot.Identity.Token,
		session.Config.Bot.Api,
	)
	Run(session)
}
