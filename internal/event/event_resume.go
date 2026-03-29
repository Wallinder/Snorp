package event

import (
	"context"
	"encoding/json"
	"log/slog"
	"snorp/internal/state"

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

func resumeConnection(ctx context.Context, conn *websocket.Conn, session *state.SessionState) {
	message, err := json.Marshal(Resume{
		Op: RESUME,
		D: ResumeData{
			Token:     session.Config.Bot.Identity.Token,
			SessionId: session.ReadyData.SessionID,
			Seq:       session.Seq,
		},
	})
	if err != nil {
		slog.Error("failed to marshal resume message", "error", err)
		return
	}
	slog.Info("Resuming connection..")

	session.Resume = false

	err = conn.Write(ctx, websocket.MessageText, message)
	if err != nil {
		slog.Error("resuming failed", "error", err)
	}
}
