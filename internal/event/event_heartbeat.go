package event

import (
	"encoding/json"
	"log"
	"menial/internal/state"
	"time"

	"golang.org/x/net/websocket"
)

type HeartbeatInterval struct {
	Interval int `json:"heartbeat_interval"`
}

type Heartbeat struct {
	Op int   `json:"op"`
	D  int64 `json:"d"`
}

func SendHeartbeat(conn *websocket.Conn, interval int, session *state.Session) {
	message, err := json.Marshal(Heartbeat{
		Op: 1,
		D:  session.GetSeq(),
	})
	if err != nil {
		log.Fatalf("Failed to marshal: %v", err)
	}
	_, err = conn.Write(message)
	if err != nil {
		log.Fatalf("Failed to send heartbeat: %v", err)
	}
	time.Sleep(time.Duration(interval) * time.Millisecond)
}
