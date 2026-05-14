package event

import (
	"context"
	"encoding/json"
	"log/slog"
	"snorp/internal/models"
	"snorp/internal/state"
)

func dispatcher(ctx context.Context, session *state.SessionState, action string, dispatchMessage json.RawMessage) {
	TotalDispatchMessages.WithLabelValues(action).Inc()

	switch action {

	case "READY":
		var readyData state.ReadyData
		if err := json.Unmarshal(dispatchMessage, &readyData); err != nil {
			slog.Info("failed to unmarshal json", "error", err)
			return
		}
		session.SetReadyData(readyData)
		session.ReadyChannel <- true

	case "GUILD_CREATE":
		var guild models.Guild
		if err := json.Unmarshal(dispatchMessage, &guild); err != nil {
			slog.Info("failed to unmarshal json", "error", err)
			return
		}

	case "INTERACTION_CREATE":
		var interaction models.Interaction
		if err := json.Unmarshal(dispatchMessage, &interaction); err != nil {
			slog.Info("failed to unmarshal json", "error", err)
			return
		}
		interactionHandler(ctx, session, interaction)
	}
}
