package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Channel struct {
	Name     string `json:"name"`
	Type     int    `json:"type"`
	Bitrate  int    `json:"bitrate"`
	Position int    `json:"position"`
	Nsfw     bool   `json:"nsfw"`
}

func CreateVoiceChannel(api, guildID, token string, client *http.Client) error {
	channel := &Channel{
		Name:     "Snorp - New Channel",
		Type:     2,
		Nsfw:     false,
		Bitrate:  16000,
		Position: 0,
	}
	body, err := json.Marshal(channel)
	if err != nil {
		return err
	}

	api = api + fmt.Sprintf("/guilds/%s/channels", guildID)

	req, err := http.NewRequest("POST", api, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header = GetHeaders(token)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	statuscode := resp.StatusCode

	if statuscode != 200 {
		return fmt.Errorf("error creating channel: %v", statuscode)
	}

	return nil
}
