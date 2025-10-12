package event

import (
	"context"
	"encoding/json"
	"log"

	"github.com/coder/websocket"
)

type Presence struct {
	Op int          `json:"op"`
	D  PresenceData `json:"d"`
}

type PresenceData struct {
	Since      int        `json:"since"`
	Activities []Activity `json:"activities"`
	Status     string     `json:"status"`
	AFK        bool       `json:"afk"`
}

type Activity struct {
	Name string `json:"name"`
	Type int    `json:"type"`
}

func UpdatePresence(ctx context.Context, conn *websocket.Conn) {
	message, err := json.Marshal(Presence{
		Op: PRESENCE_UPDATE,
		D: PresenceData{
			Since: 0,
			Activities: []Activity{
				{
					Name: "🥜Jerkmate Ranked🥜",
					Type: 5,
				},
			},
			Status: "online",
			AFK:    false,
		},
	})
	if err != nil {
		log.Printf("Failed to unmarshal presence: %v\n", err)
	}
	log.Println("Updating presence..")

	err = conn.Write(ctx, websocket.MessageText, message)
	if err != nil {
		log.Printf("Updating presence failed: %v\n", err)
	}
}
