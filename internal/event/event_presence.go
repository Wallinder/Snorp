package event

import (
	"context"
	"encoding/json"
	"log"

	"github.com/coder/websocket"
)

type Presence struct {
	Op int
	D  PresenceData
}

type PresenceData struct {
	Since      int        // Unix time (in milliseconds) of when the client went idle, or null if the client is not idle
	Activities []Activity // array of activity objects	User's activities
	Status     string     // User's new status
	AFK        bool       // Whether or not the client is afk
}

type Activity struct {
	Name string
	Type int
	// add missing
	// https://discord.com/developers/docs/events/gateway-events#activity-object
}

func UpdatePresence(ctx context.Context, conn *websocket.Conn) {
	message, err := json.Marshal(Presence{
		Op: PRESENCE_UPDATE,
		D: PresenceData{
			Since:  0,
			Status: "Gooning",
			AFK:    false,
			Activities: []Activity{
				{Name: "coding in Go 🐹", Type: 0},
			},
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
