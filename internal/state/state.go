package state

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"snorp/pkg/discord"
	"sync"
	"time"

	"github.com/coder/websocket"
)

type SessionState struct {
	StartTime    time.Time
	Metadata     discord.Metadata
	ReadyData    discord.ReadyData
	Config       *Config
	Connection   *Connection
	WsConn       *websocket.Conn
	Client       *http.Client
	CommandsDir  string
	Commands     []discord.ApplicationCommand
	ReadyChannel chan bool
	Status       Status
}

type Status struct {
	Ready   bool
	Healthy bool
}

type Connection struct {
	Mu     sync.RWMutex
	Seq    int64
	Resume bool
}

func NewState() *SessionState {
	state := newDefaultState()
	state.onReady()

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
		Client: newHttpClient(),
		Connection: &Connection{
			Resume: false,
		},
		CommandsDir: "./commands",
		StartTime:   time.Now(),
		Status: Status{
			Ready: false,
		},
		ReadyChannel: make(chan bool),
	}
}

func newHttpClient() *http.Client {
	return &http.Client{
		CheckRedirect: nil,
		Timeout:       5 * time.Second,
	}
}

func (s *SessionState) NewDiscordRequest(method string, uri string, body io.Reader) (*http.Response, error) {
	url := s.Config.Bot.Api + "/v" + s.Config.Bot.ApiVersion + uri

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	slog.Info("client request", "method", req.Method, "url", req.URL.Path)

	TotalClientHttpRequests.WithLabelValues(req.Method, req.URL.Path).Inc()

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bot %s", s.Config.Bot.Identity.Token))
	req.Header.Set("User-Agent", "DiscordBot (https://github.com/Wallinder/Snorp)")

	return s.Client.Do(req)
}

func (s *SessionState) setMetadata() {
	response, err := s.NewDiscordRequest("GET", "/gateway/bot", nil)
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

func (s *SessionState) onReady() {
	go func() {
		<-s.ReadyChannel
		s.Status.Ready = true
		s.setCommands()
	}()
}

func (s *SessionState) SetReadyData(readyData discord.ReadyData) {
	s.ReadyData = readyData
}

func (s *SessionState) SetConnection(conn *websocket.Conn) {
	s.WsConn = conn
}

func (s *SessionState) isReady() bool {
	return s.Status.Ready
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
