package event

import (
	"context"
	"encoding/json"
	"fmt"
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
		Op: 1,
		D:  seq,
	})
	if err != nil {
		log.Fatalf("Failed to marshal: %v", err)
	}
	fmt.Println(string(message))
	err = conn.Write(ctx, websocket.MessageText, message)

	if err != nil {
		log.Fatalf("Failed to send heartbeat: %v", err)
	}
}
