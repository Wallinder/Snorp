package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"snorp/internal/state"
)

type Channel struct {
	Name     string `json:"name"`
	Type     int    `json:"type"`
	Bitrate  int    `json:"bitrate"`
	Position int    `json:"position"`
	Nsfw     bool   `json:"nsfw"`
}

func CreateVoiceChannel(session *state.SessionState, guildID string, vcName string) {
	channel := &Channel{
		Name:     vcName,
		Type:     2,
		Nsfw:     false,
		Bitrate:  16000,
		Position: 0,
	}

	body, err := json.Marshal(channel)
	if err != nil {
		log.Printf("Error creating channel: %s\n", err)
	}

	request := state.HttpRequest{
		Method: "GET",
		Uri:    fmt.Sprintf("/guilds/%s/channels", guildID),
		Body:   bytes.NewBuffer(body),
	}

	_, err = session.SendRequest(request)
	if err != nil {
		log.Printf("Error creating channel: %s\n", err)
	}
}
