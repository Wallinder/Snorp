package event

import (
	"context"
	"encoding/json"
	"log"
	"snorp/config"

	"github.com/coder/websocket"
)

type PresenceUpdate struct {
	Op int                     `json:"op"`
	D  config.IdentityPresence `json:"d"`
}

func UpdatePresence(ctx context.Context, conn *websocket.Conn, presence PresenceUpdate) {
	message, err := json.Marshal(presence)
	if err != nil {
		log.Printf("Failed to unmarshal presence: %v\n", err)
	}
	log.Println("Updating presence..")

	err = conn.Write(ctx, websocket.MessageText, message)
	if err != nil {
		log.Printf("Updating presence failed: %v\n", err)
	}
}
