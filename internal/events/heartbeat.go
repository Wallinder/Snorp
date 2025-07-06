package event

import (
	"encoding/json"
	"log"
	"menial/internal/config"
	"time"

	"github.com/gorilla/websocket"
)

func SendHeartbeat(conn *websocket.Conn, interval int, HearbeatRequest config.Heartbeat) {
	message, err := json.Marshal(HearbeatRequest)
	if err != nil {
		log.Fatal("Failed to marshal", err)
	}
	err = conn.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		log.Fatal("Heartbeat failed", err)
	}
	time.Sleep(time.Duration(interval) * time.Second)
}
