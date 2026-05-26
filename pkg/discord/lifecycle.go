package discord

import (
	"context"
	"encoding/json"

	"github.com/coder/websocket"
)

type Metadata struct {
	Url               string            `json:"url"`
	Shards            int               `json:"shards"`
	SessionStartLimit SessionStartLimit `json:"session_start_limit"`
}

type SessionStartLimit struct {
	Total          int `json:"total"`
	Remaining      int `json:"remaining"`
	ResetAfter     int `json:"reset_after"`
	MaxConcurrency int `json:"max_concurrency"`
}

type ReadyData struct {
	V                    int                `json:"v"`
	UserSettings         any                `json:"user_settings"`
	User                 User               `json:"user"`
	SessionType          string             `json:"session_type"`
	SessionID            string             `json:"session_id"`
	ResumeGatewayURL     string             `json:"resume_gateway_url"`
	Relationships        any                `json:"relationships"`
	PrivateChannels      any                `json:"private_channels"`
	Presences            any                `json:"presences"`
	Guilds               []UnavailableGuild `json:"guilds"`
	GuildJoinRequests    any                `json:"guild_join_requests"`
	GeoOrderedRtcRegions []string           `json:"geo_ordered_rtc_regions"`
	GameRelationships    any                `json:"game_relationships"`
	Auth                 any                `json:"auth"`
	Application          Application        `json:"application"`
}

type Application struct {
	ID    string `json:"id"`
	Flags int    `json:"flags"`
}

type Interval struct {
	Heartbeat int `json:"heartbeat_interval"`
}

type Heartbeat struct {
	Op int   `json:"op"`
	D  int64 `json:"d"`
}

// Wrapper for heartbeat send
func SendHeartbeat(ctx context.Context, conn *websocket.Conn, seq int64) error {
	heartbeat := Heartbeat{
		Op: HEARTBEAT,
		D:  seq,
	}
	return heartbeat.Send(ctx, conn, seq)
}

func (h *Heartbeat) Send(ctx context.Context, conn *websocket.Conn, seq int64) error {
	message, err := json.Marshal(h)
	if err != nil {
		return err
	}
	return conn.Write(ctx, websocket.MessageText, message)
}

type Identify struct {
	Op int       `json:"op"`
	D  *Identity `json:"d"`
}

type Identity struct {
	Token          string             `json:"token"`
	Compress       bool               `json:"compress"`
	LargeThreshold int                `json:"large_threshold"`
	Intents        int64              `json:"intents"`
	Properties     IdentityProperties `json:"properties"`
	Presence       Presence           `json:"presence"`
	Shards         []int              `json:"shards,omitempty"`
}

type IdentityProperties struct {
	Os      string `json:"os"`
	Browser string `json:"browser"`
	Device  string `json:"device"`
}

func (i *Identity) Send(ctx context.Context, conn *websocket.Conn) error {
	message, err := json.Marshal(Identify{
		Op: IDENTIFY,
		D:  i,
	})
	if err != nil {
		return err
	}
	return conn.Write(ctx, websocket.MessageText, message)
}

type Resume struct {
	Op int        `json:"op"`
	D  ResumeData `json:"d"`
}

type ResumeData struct {
	Token     string `json:"token"`
	SessionID string `json:"session_id"`
	Seq       int64  `json:"seq"`
}

// Resume wrapper
func SendResume(ctx context.Context, conn *websocket.Conn, token string, sessionID string, seq int64) error {
	resume := Resume{
		Op: RESUME,
		D: ResumeData{
			Token:     token,
			SessionID: sessionID,
			Seq:       seq,
		},
	}
	return resume.Send(ctx, conn)
}

func (r *Resume) Send(ctx context.Context, conn *websocket.Conn) error {
	message, err := json.Marshal(r)
	if err != nil {
		return err
	}
	return conn.Write(ctx, websocket.MessageText, message)
}

// Presence
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
