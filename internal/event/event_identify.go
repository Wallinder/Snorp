package event

import (
	"context"
	"encoding/json"
	"log"
	"menial/config"

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

func SendIdentify(conn *websocket.Conn, conf config.Identity, token string) {
	message, err := json.Marshal(Identify{
		Op: 2,
		D: IdentifyData{
			Token:   token,
			Intents: conf.Intents,
			Properties: IdentifyProperties{
				Os:      conf.Properties.Os,
				Browser: conf.Properties.Browser,
				Device:  conf.Properties.Device,
			},
		},
	})
	if err != nil {
		log.Fatalf("Failed to unmarshal identity: %v", err)
	}
	log.Println("Identifying..")
	err = conn.Write(context.TODO(), websocket.MessageText, message)
	if err != nil {
		log.Fatalf("Identity failed: %v", err)
	}
}
