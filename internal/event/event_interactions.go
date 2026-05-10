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

func InteractionHandler(ctx context.Context, session *state.SessionState, interaction models.Interaction) {
	switch interaction.Data.Name {
	case "politiloggen":
		for _, option := range interaction.Data.Options {
			switch option.Name {
			case "nyeste":
				msg, _ := politiloggen.GetLastMessage()
				interactionCallback(session, interaction, msg.Data.Text)
			}
		}
	}
}

type Callback struct {
	Type int          `json:"type"`
	Data CallbackData `json:"data"`
}

type CallbackData struct {
	Content string `json:"content"`
}

func interactionCallback(session *state.SessionState, interaction models.Interaction, msg string) {
	uri := "/interactions/" + interaction.ID + "/" + interaction.Token + "/callback"

	callback := Callback{
		Type: 4,
		Data: CallbackData{
			Content: msg,
		},
	}

	data, err := json.Marshal(callback)
	if err != nil {
		callback.Data.Content = "failed to fetch data"
	}

	_, err = session.NewRequest("POST", uri, bytes.NewReader(data))
	if err != nil {
		slog.Error("callback failed", "error", err)
	}
}
