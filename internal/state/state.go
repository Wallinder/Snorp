package state

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"os"
	"snorp/config"
	"sync"
	"time"

	"github.com/coder/websocket"
	"github.com/prometheus/client_golang/prometheus"
)

type SessionState struct {
	Mu            sync.Mutex
	StartTime     time.Time
	Seq           int64
	Metadata      Metadata
	ReadyData     ReadyData
	Resume        bool
	Config        *config.Config
	Conn          *websocket.Conn
	Client        *http.Client
	Metrics       *Metrics
	GlobalHeaders map[string][]string
	Messages      chan []byte
	MaxRetries    int
}

type Metrics struct {
	Uri                   string
	Port                  int
	TotalMessages         *prometheus.CounterVec
	TotalDispatchMessages *prometheus.CounterVec
	TotalHttpRequests     *prometheus.CounterVec
	TotalDisconnects      prometheus.Counter
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
	state.newHttpClient()
	state.setMetadata()
	if len(state.Config.Bot.Identity.Shards) == 0 {
		shards := []int{0, 1}
		slog.Info("using default sharding", "shards", shards)
		state.Config.Bot.Identity.Shards = shards
	}
	slog.Info("recommended sharding", "shards", state.Metadata.Shards)
	return state
}

func newDefaultState() *SessionState {
	return &SessionState{
		Config:     config.NewConfig(),
		Resume:     false,
		Messages:   make(chan []byte),
		MaxRetries: 3,
		Metrics: &Metrics{
			Uri:  "/metrics",
			Port: 8080,
		},
		StartTime: time.Now(),
	}
}

func (s *SessionState) setMetadata() {
	response, err := s.NewDiscordRequest("GET", "/gateway/bot", nil)
	if err != nil {
		slog.Error("unable to update metadata", "error", err)
		os.Exit(1)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		slog.Error("unable to update metadata", "error", err)
		os.Exit(1)
	}

	if err = json.Unmarshal(body, &s.Metadata); err != nil {
		slog.Error("unable to update metadata", "error", err)
		os.Exit(1)
	}
}
