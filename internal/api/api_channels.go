package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"snorp/internal/state"
)

type Channel struct {
	Name     string `json:"name"`
	Type     int    `json:"type"`
	Bitrate  int    `json:"bitrate"`
	Position int    `json:"position"`
	Nsfw     bool   `json:"nsfw"`
}

func CreateVoiceChannel(session *state.SessionState, guildID string, guildName string) {
	log.Printf("Creating VC in %s\n", guildName)
	body, err := json.Marshal(
		&Channel{
			Name:     "♻️New Voice Channel♻️",
			Type:     2,
			Nsfw:     false,
			Bitrate:  16000,
			Position: 0,
		},
	)
	if err != nil {
		log.Printf("Error creating channel: %s\n", err)
		return
	}

	api := session.Config.Bot.Api + fmt.Sprintf("/guilds/%s/channels", guildID)

	request, err := http.NewRequest("POST", api, bytes.NewBuffer(body))
	if err != nil {
		log.Printf("Error creating channel: %s\n", err)
		return
	}

	request.Header = session.GlobalHeaders

	response, err := session.Client.Do(request)
	if err != nil {
		log.Printf("Error creating channel: %s\n", err)
		return
	}
	defer response.Body.Close()

	statuscode := response.StatusCode

	if statuscode == 403 {
		log.Printf("Error creating, missing permissions: %d\n", statuscode)
		return
	}
	if statuscode != 200 && statuscode != 201 {
		log.Printf("Error creating channel: %d\n", statuscode)
		return
	}
}
