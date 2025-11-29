package state

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"snorp/config"
	"sync"
	"time"

	"github.com/coder/websocket"
	"github.com/prometheus/client_golang/prometheus"
	"gorm.io/gorm"
)

type SessionState struct {
	Mu            sync.Mutex
	StartTime     time.Time
	Seq           int64
	Metadata      Metadata
	ReadyData     ReadyData
	Resume        bool
	DB            *gorm.DB
	Config        config.Config
	Conn          *websocket.Conn
	Client        *http.Client
	Metrics       *Metrics
	GlobalHeaders map[string][]string
	Messages      chan []byte
	MaxRetries    int
	Jobs          Jobs
}

type Metrics struct {
	Uri                   string
	Port                  int
	TotalMessages         *prometheus.CounterVec
	TotalDispatchMessages *prometheus.CounterVec
	TotalHttpRequests     *prometheus.CounterVec
	TotalDisconnects      prometheus.Counter
}

type Jobs struct {
	Welcome    map[string]string
	SteamNews  map[string]bool
	SteamSales map[string]bool
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

func (session *SessionState) InitHttpClient() *http.Client {
	session.Client = &http.Client{
		CheckRedirect: nil,
		Timeout:       time.Duration(10 * time.Second),
	}
	session.Client.Transport = &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    10 * time.Second,
		DisableCompression: true,
	}
	session.GlobalHeaders = map[string][]string{
		"Content-Type":  {"application/json"},
		"User-Agent":    {"DiscordBot (https://github.com/Wallinder/Snorp)"},
		"Authorization": {fmt.Sprintf("Bot %s", session.Config.Bot.Identity.Token)},
	}
	return session.Client
}

func (session *SessionState) UpdateMetadata() {
	request := HttpRequest{
		Method: "GET",
		Uri:    "/gateway/bot",
		Body:   nil,
	}

	response, err := session.SendRequest(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var metadata Metadata

	err = json.Unmarshal(body, &metadata)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Recommended amount of shards: %d\n", metadata.Shards)

	if len(session.Config.Bot.Identity.Shards) == 0 {
		shards := []int{0, 1}
		log.Printf("No shards in config, using default value: %d", shards)

		session.Config.Bot.Identity.Shards = shards
	}

	session.Metadata = metadata
}
