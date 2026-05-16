package discord

import (
	"context"
	"encoding/json"

	"github.com/coder/websocket"
)

type Presence struct {
	User                 User         `json:"user"`
	Status               string       `json:"status,omitempty"`
	Since                int          `json:"since,omitempty"`
	ProcessedAtTimestamp int64        `json:"processed_at_timestamp,omitempty"`
	ClientStatus         ClientStatus `json:"client_status"`
	Activities           []*Activity  `json:"activities,omitempty"`
	AFK                  bool         `json:"afk,omitempty"`
}

type Activity struct {
	Name string `json:"name"`
	Type int    `json:"type"`
}

type ClientStatus struct {
	Desktop string `json:"desktop"`
}

type PresenceUpdate struct {
	Op int       `json:"op"`
	D  *Presence `json:"d"`
}

func (p *Presence) Update(ctx context.Context, conn *websocket.Conn, presence PresenceUpdate) error {
	message, err := json.Marshal(PresenceUpdate{Op: PRESENCE_UPDATE, D: p})
	if err != nil {
		return err
	}
	return conn.Write(ctx, websocket.MessageText, message)
}
