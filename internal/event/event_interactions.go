package event

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"snorp/internal/models"
	"snorp/internal/state"
	"snorp/pkg/politiloggen"
)

func interactionHandler(_ context.Context, session *state.SessionState, interaction models.Interaction) {
	switch interaction.Data.Name {
	case "politiloggen":
		for _, option := range interaction.Data.Options {
			switch option.Name {
			case "nyeste":
				msg, _ := politiloggen.GetLastMessage()
				interactionCallback(session, interaction, models.InteractionCallback{
					Type: models.CallbackChannelMessageWithSource,
					Data: models.InteractionCallbackData{
						Content: msg.Data.Text,
					},
				})
			}
		}
	}
}

func interactionCallback(session *state.SessionState, interaction models.Interaction, callback models.InteractionCallback) {
	uri := "/interactions/" + interaction.ID + "/" + interaction.Token + "/callback"

	data, err := json.Marshal(callback)
	if err != nil {
		slog.Error("callback marshal", "error", err)
	}

	_, err = session.NewRequest("POST", uri, bytes.NewReader(data))
	if err != nil {
		slog.Error("callback failed", "error", err)
	}
}
