package event

import (
	"context"
	"encoding/json"
	"log"
	"snorp/internal/state"

	"github.com/coder/websocket"
)

type Identify struct {
	Op int          `json:"op"`
	D  IdentifyData `json:"d"`
}

type IdentifyData struct {
	Token      string             `json:"token"`
	Intents    int64              `json:"intents"`
	Properties IdentifyProperties `json:"properties"`
}

type IdentifyProperties struct {
	Os      string `json:"os"`
	Browser string `json:"browser"`
	Device  string `json:"device"`
}

func SendIdentify(ctx context.Context, conn *websocket.Conn, session *state.SessionState) {
	message, err := json.Marshal(Identify{
		Op: 2,
		D: IdentifyData{
			Token:   session.Config.Bot.Token,
			Intents: session.Config.Bot.Identity.Intents,
			Properties: IdentifyProperties{
				Os:      session.Config.Bot.Identity.Properties.Os,
				Browser: session.Config.Bot.Identity.Properties.Browser,
				Device:  session.Config.Bot.Identity.Properties.Device,
			},
		},
	})
	if err != nil {
		log.Fatalf("Failed to unmarshal identity: %v", err)
	}
	log.Println("Identifying..")

	err = conn.Write(ctx, websocket.MessageText, message)
	if err != nil {
		log.Fatalf("Identity failed: %v", err)
	}
}
