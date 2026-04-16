package event

import (
	"context"
	"encoding/json"
	"log/slog"
	"snorp/internal/models"
	"snorp/internal/state"
)

func Dispatcher(ctx context.Context, session *state.SessionState, action string, dispatchMessage json.RawMessage) {
	TotalDispatchMessages.WithLabelValues(action).Inc()
	slog.Info("event", "type", "DISPATCH", "action", action)

	switch action {

	case "READY":
		var readyData state.ReadyData
		if err := json.Unmarshal(dispatchMessage, &readyData); err != nil {
			slog.Info("failed to unmarshal json", "error", err)
		}
		session.SetReadyData(readyData)

	case "RESUMED":

	case "GUILD_CREATE":
		var guild models.Guild
		if err := json.Unmarshal(dispatchMessage, &guild); err != nil {
			slog.Info("failed to unmarshal json", "error", err)
		}

	default:
		//fmt.Println(string(dispatchMessage))
	}
}
