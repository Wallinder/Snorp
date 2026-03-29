package event

import (
	"context"
	"encoding/json"
	"log/slog"
	"snorp/config"

	"github.com/coder/websocket"
)

type Identify struct {
	Op int             `json:"op"`
	D  config.Identity `json:"d"`
}

func identify(ctx context.Context, conn *websocket.Conn, identity config.Identity) {
	message, err := json.Marshal(Identify{
		Op: IDENTIFY,
		D:  identity,
	})
	if err != nil {
		slog.Error("failed to unmarshal identity", "error", err)
		return
	}
	slog.Info("identifying..")

	err = conn.Write(ctx, websocket.MessageText, message)
	if err != nil {
		slog.Error("identity failed", "error", err)
	}
}
