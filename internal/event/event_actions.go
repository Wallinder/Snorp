package event

import (
	"context"
	"encoding/json"
	"log/slog"
	"snorp/internal/models"
	"snorp/internal/state"

	"github.com/coder/websocket"
)

// HEARTBEAT

type Interval struct {
	Heartbeat int `json:"heartbeat_interval"`
}

type Heartbeat struct {
	Op int   `json:"op"`
	D  int64 `json:"d"`
}

func sendHeartbeat(ctx context.Context, conn *websocket.Conn, seq int64) {
	message, err := json.Marshal(Heartbeat{
		Op: HEARTBEAT,
		D:  seq,
	})
	if err != nil {
		slog.Error("failed to marshal heartbeat", "error", err)
	}
	err = conn.Write(ctx, websocket.MessageText, message)
	if err != nil {
		slog.Error("failed to send heartbeat", "error", err)
	}
}

// IDENTIFY

type Identify struct {
	Op int            `json:"op"`
	D  state.Identity `json:"d"`
}

func identify(ctx context.Context, conn *websocket.Conn, identity state.Identity) {
	message, err := json.Marshal(Identify{
		Op: IDENTIFY,
		D:  identity,
	})
	if err != nil {
		slog.Error("failed to unmarshal identity", "error", err)
		return
	}
	err = conn.Write(ctx, websocket.MessageText, message)
	if err != nil {
		slog.Error("identity failed", "error", err)
	}
}

// PRESENCE

type PresenceUpdate struct {
	Op int             `json:"op"`
	D  models.Presence `json:"d"`
}

func updatePresence(ctx context.Context, conn *websocket.Conn, presence PresenceUpdate) {
	message, err := json.Marshal(presence)
	if err != nil {
		slog.Error("failed to unmarshal presence", "error", err)
		return
	}
	err = conn.Write(ctx, websocket.MessageText, message)
	if err != nil {
		slog.Error("updating presence failed", "error", err)
	}
}

// RESUME

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

	session.SetResume(false)

	err = conn.Write(ctx, websocket.MessageText, message)
	if err != nil {
		slog.Error("resuming failed", "error", err)
	}
}
