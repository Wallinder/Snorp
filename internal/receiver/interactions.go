package receiver

import (
	"bytes"
	"context"
	"encoding/json"
	"snorp/internal/state"
	"snorp/pkg/discord"
)

func interactionHandler(_ context.Context, session *state.SessionState, interaction discord.Interaction) {
	switch interaction.Data.Name {
	}
}

func interactionCallback(session *state.SessionState, interaction discord.Interaction, callback discord.InteractionCallback) {
	uri := "/interactions/" + interaction.ID + "/" + interaction.Token + "/callback"

	data, err := json.Marshal(callback)
	if err != nil {
		session.ErrorChan <- state.SessionError{Origin: "interaction", Err: err, Fatal: false}
	}

	_, err = session.Discord.NewDiscordRequest("POST", uri, bytes.NewReader(data))
	if err != nil {
		session.ErrorChan <- state.SessionError{Origin: "interaction", Err: err, Fatal: false}
	}
}
