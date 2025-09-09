package event

import (
	"context"
	"encoding/json"
	"log"
	"menial/internal/state"

	"github.com/coder/websocket"
)

type Resume struct {
	Op int        `json:"op"`
	D  ResumeData `json:"d"`
}

type ResumeData struct {
	Token     string `json:"token"`
	SessionId string `json:"session_id"`
	Seq       int64  `json:"seq"`
}

func ResumeConnection(ctx context.Context, conn *websocket.Conn, session *state.SessionState) {
	message, err := json.Marshal(Resume{
		Op: 6,
		D: ResumeData{
			Token:     session.Config.Bot.Token,
			SessionId: session.ReadyData.SessionID,
			Seq:       session.Seq,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Resuming connection..")

	err = conn.Write(ctx, websocket.MessageText, message)
	if err != nil {
		log.Printf("Resuming failed: %s\n", err)
		session.Resume = false
	}
}
