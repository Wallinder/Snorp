package discord

import (
	"context"
	"encoding/json"

	"github.com/coder/websocket"
)

const (
	DISPATCH                  = 0  //  Receive       An event was dispatched.
	HEARTBEAT                 = 1  //  Send/Receive  Fired periodically by the client to keep the connection alive.
	IDENTIFY                  = 2  //  Send	         Starts a new session during the initial handshake.
	PRESENCE_UPDATE           = 3  //  Send          Update the client's presence.
	VOICE_STATE_UPDATE        = 4  //  Send          Used to join/leave or move between voice channels.
	RESUME                    = 6  //  Send          Resume a previous session that was disconnected.
	RECONNECT                 = 7  //  Receive       You should attempt to reconnect and resume immediately.
	REQUEST_GUILD_MEMBERS     = 8  //  Send          Request information about offline guild members in a large guild.
	INVALID_SESSION           = 9  //  Receive       The session has been invalidated. You should reconnect and identify/resume accordingly.
	HELLO                     = 10 //  Receive       Sent immediately after connecting, contains the heartbeat_interval to use.
	HEARTBEAT_ACK             = 11 //  Receive       Sent in response to receiving a heartbeat to acknowledge that it has been received.
	REQUEST_SOUNDBOARD_SOUNDS = 31 //  Send          Request information about soundboard sounds in a set of guilds.
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
