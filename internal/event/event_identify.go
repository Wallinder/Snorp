package event

import (
	"context"
	"encoding/json"
	"log"
	"snorp/config"

	"github.com/coder/websocket"
)

type Identify struct {
	Op int             `json:"op"`
	D  config.Identity `json:"d"`
}

func SendIdentify(ctx context.Context, conn *websocket.Conn, identity config.Identity) {
	message, err := json.Marshal(Identify{
		Op: 2,
		D:  identity,
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
