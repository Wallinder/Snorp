package state

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"snorp/config"
	"time"

	"github.com/coder/websocket"
)

type SessionState struct {
	Seq        int64
	Metadata   Metadata
	ReadyData  ReadyData
	Resume     bool
	Config     config.Config
	Conn       *websocket.Conn
	Client     *http.Client
	Messages   chan []byte
	MaxRetries int
}

type ReadyData struct {
	V                    int         `json:"v"`
	UserSettings         any         `json:"user_settings"`
	User                 User        `json:"user"`
	SessionType          string      `json:"session_type"`
	SessionID            string      `json:"session_id"`
	ResumeGatewayURL     string      `json:"resume_gateway_url"`
	Relationships        any         `json:"relationships"`
	PrivateChannels      any         `json:"private_channels"`
	Presences            any         `json:"presences"`
	Guilds               any         `json:"guilds"`
	GuildJoinRequests    any         `json:"guild_join_requests"`
	GeoOrderedRtcRegions []string    `json:"geo_ordered_rtc_regions"`
	GameRelationships    any         `json:"game_relationships"`
	Auth                 any         `json:"auth"`
	Application          Application `json:"application"`
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

func (s *SessionState) InitHttpClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    10 * time.Second,
			DisableCompression: true,
		},
		CheckRedirect: nil,
		Timeout:       time.Duration(10 * time.Second),
	}
}

func (s *SessionState) UpdateMetadata(token string, api string) {
	client := s.InitHttpClient()

	gateway := api + "/gateway/bot"

	req, err := http.NewRequest("GET", gateway, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bot %s", token))

	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	var metadata *Metadata
	err = json.Unmarshal(body, &metadata)
	if err != nil {
		log.Fatal(err)
	}
	s.Metadata = *metadata
}
