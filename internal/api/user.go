package api

import (
	"encoding/json"
	"fmt"
	"io"
	"snorp/internal/state"
)

type User struct {
	ID                   string                `json:"id"`
	Username             string                `json:"username,omitzero"`
	Discriminator        string                `json:"discriminator,omitzero"`
	GlobalName           string                `json:"global_name,omitzero"`
	Avatar               string                `json:"avatar,omitzero"`
	Bot                  bool                  `json:"bot,omitzero"`
	Email                string                `json:"email,omitzero"`
	System               bool                  `json:"system,omitzero"`
	MfaEnabled           bool                  `json:"mfa_enabled,omitzero"`
	PublicFlags          int                   `json:"public_flags,omitzero"`
	PrimaryGuild         GuildUserPrimaryGuild `json:"primary_guild,omitzero"`
	DisplayNameStyles    any                   `json:"display_name_styles,omitzero"`
	DisplayName          string                `json:"display_name,omitzero"`
	Collectibles         any                   `json:"collectibles,omitzero"`
	AvatarDecorationData any                   `json:"avatar_decoration_data,omitzero"`
}

func GetUser(session *state.SessionState, userID string) (User, error) {
	request := state.HttpRequest{
		Method: "GET",
		Uri:    fmt.Sprintf("/users/%s", userID),
		Body:   nil,
	}

	var user User

	response, err := session.SendRequest(request)
	if err != nil {
		return user, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return user, err
	}
	defer response.Body.Close()

	err = json.Unmarshal(body, &user)
	if err != nil {
		return user, err
	}

	return user, nil
}
