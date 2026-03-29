package dispatcher

import (
	"context"
	"encoding/json"
	"log/slog"
	"snorp/internal/state"
)

func Actions(ctx context.Context, session *state.SessionState, action string, dispatchMessage json.RawMessage) {
	TotalDispatchMessages.WithLabelValues(action).Inc()

	switch action {

	case "READY":
		slog.Info("handshake complete..")
		var readyData state.ReadyData
		if err := json.Unmarshal(dispatchMessage, &readyData); err != nil {
			slog.Info("failed to unmarshal json", "error", err)
		}
		session.ReadyData = readyData

	case "GUILD_MEMBER_ADD":

	case "GUILD_CREATE":

	case "GUILD_DELETE":

	case "GUILD_UPDATE":

	case "CHANNEL_CREATE":

	case "CHANNEL_DELETE":

	case "CHANNEL_UPDATE":

	case "MESSAGE_CREATE":

	case "INTERACTION_CREATE":

	case "RESUMED":
		slog.Info("connection resumed")
	}
}
