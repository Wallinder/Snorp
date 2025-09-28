package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"snorp/internal/state"
)

func (gc *GuildChannels) CreateChannel(session *state.SessionState, guildID string) (*http.Response, error) {
	body, err := json.Marshal(gc)
	if err != nil {
		return nil, err
	}

	request := state.HttpRequest{
		Method: "POST",
		Uri:    fmt.Sprintf("/guilds/%s/channels", guildID),
		Body:   bytes.NewBuffer(body),
	}

	response, err := session.SendRequest(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
