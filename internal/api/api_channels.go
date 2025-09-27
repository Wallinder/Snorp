package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"snorp/internal/state"
)

type NewChannel struct {
	Name        string                     `json:"name"`
	Type        int                        `json:"type"`
	Position    int                        `json:"position"`
	Permissions []GuildChannelsPermissions `json:"permission_overwrites,omitzero"`
	Bitrate     int                        `json:"bitrate,omitzero"`
	Nsfw        bool                       `json:"nsfw,omitzero"`
	ParentID    string                     `json:"parent_id,omitzero"`
}

func (nc *NewChannel) CreateChannel(session *state.SessionState, guildID string) (*http.Response, error) {
	body, err := json.Marshal(nc)
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
