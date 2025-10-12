package event

import (
	"context"
	"encoding/json"
	"log"

	"github.com/coder/websocket"
)

type HeartbeatInterval struct {
	Interval int `json:"heartbeat_interval"`
}

type Heartbeat struct {
	Op int   `json:"op"`
	D  int64 `json:"d"`
}

func SendHeartbeat(ctx context.Context, conn *websocket.Conn, seq int64) {
	message, err := json.Marshal(Heartbeat{
		Op: HEARTBEAT,
		D:  seq,
	})
	if err != nil {
		log.Fatalf("Failed to marshal: %v", err)
	}
	err = conn.Write(ctx, websocket.MessageText, message)
	if err != nil {
		log.Fatalf("Failed to send heartbeat: %v", err)
	}
}
