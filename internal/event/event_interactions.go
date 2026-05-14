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
	callback := newCallback()
	switch interaction.Data.Name {
	case "politiloggen":
		for _, option := range interaction.Data.Options {
			switch option.Name {
			case "nyeste":
				msg, _ := politiloggen.GetLastMessage()
				callback.Type = models.CallbackChannelMessageWithSource
				callback.Data.Content = msg.Data.Text
				interactionCallback(session, interaction, callback)
			}
		}
	}
}

func newCallback() models.InteractionCallback {
	return models.InteractionCallback{
		Data: models.InteractionCallbackData{},
	}
}

func interactionCallback(session *state.SessionState, interaction models.Interaction, callback models.InteractionCallback) {
	uri := "/interactions/" + interaction.ID + "/" + interaction.Token + "/callback"

	data, err := json.Marshal(callback)
	if err != nil {
		slog.Error("callback failed", "error", err)
	}

	_, err = session.NewRequest("POST", uri, bytes.NewReader(data))
	if err != nil {
		slog.Error("callback failed", "error", err)
	}
}
