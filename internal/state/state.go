package state

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"os"
	"snorp/internal/models"
	"sync"
	"time"

	"github.com/coder/websocket"
)

type SessionState struct {
	StartTime  time.Time
	Metadata   Metadata
	ReadyData  ReadyData
	Config     *Config
	Connection *Connection
	WsConn     *websocket.Conn
	Client     *http.Client
	Commands   []*models.ApplicationCommand
}

type Connection struct {
	Mu     sync.RWMutex
	Seq    int64
	Resume bool
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

type UnavailableGuild struct {
	ID          string
	Unavailable bool
}

type User struct {
	Verified      bool   `json:"verified"`
	Username      string `json:"username"`
	PrimaryGuild  any    `json:"primary_guild"`
	MfaEnabled    bool   `json:"mfa_enabled"`
	ID            string `json:"id"`
	GlobalName    any    `json:"global_name"`
	Flags         int    `json:"flags"`
	Email         any    `json:"email"`
	Discriminator string `json:"discriminator"`
	Clan          any    `json:"clan"`
	Bot           bool   `json:"bot"`
	Avatar        any    `json:"avatar"`
}

type Application struct {
	ID    string `json:"id"`
	Flags int    `json:"flags"`
}

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

func NewState() *SessionState {
	state := newDefaultState()
	state.Client = newHttpClient(state.Config.Bot.Identity.Token)

	state.setMetadata()
	if len(state.Config.Bot.Identity.Shards) == 0 {
		shards := []int{0, 1}
		slog.Info("using default sharding", "shards", shards)
		state.Config.Bot.Identity.Shards = shards
	}
	slog.Info("recommended by discord", "shards", state.Metadata.Shards)
	return state
}

func newDefaultState() *SessionState {
	return &SessionState{
		Config: NewConfig(),
		Connection: &Connection{
			Resume: false,
		},
		StartTime: time.Now(),
	}
}

func (s *SessionState) setMetadata() {
	response, err := s.NewRequest("GET", "/gateway/bot", nil)
	if err != nil {
		LogAndExit("unable to send discord request", err, 1)
	}
	defer response.Body.Close()

	statusCode := response.StatusCode
	if statusCode < 200 || statusCode >= 300 {
		LogAndExit("received bad statuscode", errors.New(response.Status), 1)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		LogAndExit("unable to read discord metadata", err, 1)
	}

	if err = json.Unmarshal(body, &s.Metadata); err != nil {
		LogAndExit("unable to update metadata", err, 1)
	}
}

func LogAndExit(msg string, err error, exitcode int) {
	slog.Error(msg, "error", err)
	os.Exit(exitcode)
}

func (s *SessionState) SetReadyData(readyData ReadyData) {
	s.ReadyData = readyData
}

func (s *SessionState) SetConnection(conn *websocket.Conn) {
	s.WsConn = conn
}

func (c *Connection) SetResume(resume bool) {
	c.Mu.Lock()
	c.Resume = resume
	c.Mu.Unlock()
}

func (c *Connection) SetSequence(seq int64) {
	c.Mu.Lock()
	c.Seq = seq
	c.Mu.Unlock()
}
