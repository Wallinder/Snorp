package event

import (
	"context"
	"encoding/json"
	"log"

	"github.com/coder/websocket"
)

type Resume struct {
	Op int
	D  ResumeData
}

type ResumeData struct {
	Token     string `json:"token"`
	SessionId string `json:"session_id"`
	Seq       int64  `json:"seq"`
}

func ResumeConnection(conn *websocket.Conn, token string, sessionId string, seq int64) {
	message, err := json.Marshal(Resume{
		Op: 6,
		D: ResumeData{
			Token:     token,
			SessionId: sessionId,
			Seq:       seq,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Resuming connection..")
	err = conn.Write(context.TODO(), websocket.MessageText, message)
	if err != nil {
		log.Fatalf("Resuming failed: %v", err)
	}
}
