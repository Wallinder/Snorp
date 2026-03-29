package event

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/coder/websocket"
)

type Interval struct {
	Heartbeat int `json:"heartbeat_interval"`
}

type Heartbeat struct {
	Op int   `json:"op"`
	D  int64 `json:"d"`
}

func sendHeartbeat(ctx context.Context, conn *websocket.Conn, seq int64) {
	message, err := json.Marshal(Heartbeat{
		Op: HEARTBEAT,
		D:  seq,
	})
	if err != nil {
		slog.Error("failed to marshal heartbeat", "error", err)
	}
	err = conn.Write(ctx, websocket.MessageText, message)
	if err != nil {
		slog.Error("failed to send heartbeat", "error", err)
	}
}
