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
		slog.Info("handshake complete", "action", action)
		var readyData state.ReadyData
		if err := json.Unmarshal(dispatchMessage, &readyData); err != nil {
			slog.Info("failed to unmarshal json", "error", err)
		}
		session.SetReadyData(readyData)

	case "RESUMED":
		slog.Info("connection resumed", "action", action)

	default:
		//fmt.Println(string(dispatchMessage))
	}
}
