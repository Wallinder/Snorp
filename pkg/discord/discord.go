package discord

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/coder/websocket"
)

type Discord struct {
	Identity     Identity
	Api          string
	ApiVersion   string
	Websocket    *Websocket
	HttpClient   *http.Client
	ReadyData    *ReadyData
	Metadata     Metadata
	Connection   *DiscordConnection
	DispatchChan chan DispatchMessage
}

type Websocket struct {
	Conn              *websocket.Conn
	MaxRetries        int
	ResetAfter        time.Duration
	ReconnectAttempts int
	LastAttempt       time.Time
	ErrorChan         chan error
}

type DiscordConnection struct {
	Mu     sync.RWMutex
	Seq    int64
	Resume bool
}

var (
	ErrMissingIdentity       = errors.New("missing identity")
	ErrBadMetadataStatusCode = errors.New("received bad statuscode when fetching metadata")
	ErrReadingMetadataBody   = errors.New("unable to read discord metadata body")
	ErrUnmarshalMetadata     = errors.New("unable to unmarshal metadata")
	ErrUnableToSendRequest   = errors.New("unable to send discord request")
)

func NewDiscord(client *http.Client, identity Identity, api string, apiVersion string) (*Discord, error) {
	if client == nil {
		client = http.DefaultClient
	}

	if api == "" || apiVersion == "" {
		return nil, fmt.Errorf("missing api or apiversion")
	}

	discord := &Discord{
		Api:        api,
		ApiVersion: apiVersion,
		Identity:   identity,
		HttpClient: client,
		Connection: &DiscordConnection{
			Resume: false,
		},
		Websocket: &Websocket{
			MaxRetries: 3,
			ResetAfter: 30 * time.Second,
			ErrorChan:  make(chan error),
		},
		DispatchChan: make(chan DispatchMessage),
	}

	if err := discord.setMetadata(); err != nil {
		return nil, err
	}

	if len(discord.Identity.Shards) == 0 {
		discord.Identity.Shards = []int{0, 1}
	}
	return discord, nil
}

func (d *Discord) NewDiscordRequest(method string, uri string, body io.Reader) (*http.Response, error) {
	url := d.Api + "/v" + d.ApiVersion + uri

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bot %s", d.Identity.Token))
	req.Header.Set("User-Agent", "DiscordBot (https://github.com/Wallinder/Snorp)")

	return d.HttpClient.Do(req)
}

func (d *Discord) setMetadata() error {
	response, err := d.NewDiscordRequest("GET", "/gateway/bot", nil)
	if err != nil {
		return ErrUnableToSendRequest
	}
	defer response.Body.Close()

	statusCode := response.StatusCode
	if statusCode < 200 || statusCode >= 300 {
		return ErrBadMetadataStatusCode
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return ErrReadingMetadataBody
	}

	if err = json.Unmarshal(body, &d.Metadata); err != nil {
		return ErrUnmarshalMetadata
	}
	return nil
}

func (dc *DiscordConnection) SetResume(resume bool) {
	dc.Mu.Lock()
	dc.Resume = resume
	dc.Mu.Unlock()
}

func (dc *DiscordConnection) SetSequence(seq int64) {
	dc.Mu.Lock()
	dc.Seq = seq
	dc.Mu.Unlock()
}

func (d *Discord) SetReadyData(readyData ReadyData) {
	d.ReadyData = &readyData
}

func (d *Discord) SetConnection(conn *websocket.Conn) {
	d.Websocket.Conn = conn
}
