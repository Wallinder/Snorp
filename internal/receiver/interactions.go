package receiver

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
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
		slog.Error("callback marshal", "error", err)
	}

	_, err = session.Discord.NewDiscordRequest("POST", uri, bytes.NewReader(data))
	if err != nil {
		slog.Error("callback failed", "error", err)
	}
}
