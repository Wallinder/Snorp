package event

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/coder/websocket"
)

type HeartbeatInterval struct {
	Interval int `json:"heartbeat_interval"`
}

type Heartbeat struct {
	Op int   `json:"op"`
	D  int64 `json:"d"`
}

func SendHeartbeat(conn *websocket.Conn, interval int, seq int64) {
	message, err := json.Marshal(Heartbeat{
		Op: 1,
		D:  seq,
	})
	if err != nil {
		log.Fatalf("Failed to marshal: %v", err)
	}
	err = conn.Write(context.TODO(), websocket.MessageText, message)
	if err != nil {
		log.Fatalf("Failed to send heartbeat: %v", err)
	}
	time.Sleep(time.Duration(interval) * time.Millisecond)
}
