package event

import (
	"context"
	"encoding/json"
	"log"
	"menial/internal/state"

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

func ResumeConnection(ctx context.Context, conn *websocket.Conn, token string, sessionState *state.SessionState) {
	message, err := json.Marshal(Resume{
		Op: 6,
		D: ResumeData{
			Token:     token,
			SessionId: sessionState.ReadyData.SessionID,
			Seq:       sessionState.Seq,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Resuming connection..")
	err = conn.Write(ctx, websocket.MessageText, message)
	if err != nil {
		log.Fatalf("Resuming failed: %v", err)
	}
}
