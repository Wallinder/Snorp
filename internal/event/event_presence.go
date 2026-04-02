package event

import (
	"context"
	"encoding/json"
	"log/slog"
	"snorp/config"

	"github.com/coder/websocket"
)

type PresenceUpdate struct {
	Op int                     `json:"op"`
	D  config.IdentityPresence `json:"d"`
}

func updatePresence(ctx context.Context, conn *websocket.Conn, presence PresenceUpdate) {
	message, err := json.Marshal(presence)
	if err != nil {
		slog.Error("failed to unmarshal presence", "error", err)
		return
	}
	err = conn.Write(ctx, websocket.MessageText, message)
	if err != nil {
		slog.Error("updating presence failed", "error", err)
	}
}
