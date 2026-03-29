package config

import (
	"encoding/json"
	"log/slog"
	"os"
)

type Config struct {
	Bot DiscordBot `json:"discord_bot"`
}

type DiscordBot struct {
	SuperuserID string   `json:"superuser_id"`
	Permissions int64    `json:"permissions"`
	Gateway     string   `json:"gateway"`
	Api         string   `json:"api"`
	ApiVersion  string   `json:"api_version"`
	Identity    Identity `json:"identity"`
}

type Identity struct {
	Token          string             `json:"token"`
	Compress       bool               `json:"compress"`
	LargeThreshold int                `json:"large_threshold"`
	Intents        int64              `json:"intents"`
	Properties     IdentityProperties `json:"properties"`
	Presence       IdentityPresence   `json:"presence"`
	Shards         []int              `json:"shards,omitempty"`
}

type IdentityProperties struct {
	Os      string `json:"os"`
	Browser string `json:"browser"`
	Device  string `json:"device"`
}

type IdentityPresence struct {
	Since      int        `json:"since"`
	Activities []Activity `json:"activities"`
	Status     string     `json:"status"`
	AFK        bool       `json:"afk"`
}

type Activity struct {
	Name string `json:"name"`
	Type int    `json:"type"`
}

func newDefaultConfig() *Config {
	return &Config{
		Bot: DiscordBot{
			SuperuserID: "216244586165698560",
			Permissions: 2031514586918128,
			Api:         "https://discord.com/api",
			ApiVersion:  "10",
			Identity: Identity{
				Properties: IdentityProperties{
					Os:      "Linux",
					Browser: "https://github.com/Wallinder/Snorp",
					Device:  "Walle-Lab",
				},
				Presence: IdentityPresence{
					Since: 0,
					Activities: []Activity{
						{
							Name: "🥜Jerkmate Ranked🥜",
							Type: 5,
						},
					},
					Status: "online",
					AFK:    false,
				},
				Token:          os.Getenv("DISCORD_TOKEN"),
				Compress:       false,
				LargeThreshold: 250,
				Intents:        130955,
			},
		},
	}
}

func (c *Config) readJsonConfig() {
	fileContent, err := os.ReadFile("config.json")
	if err != nil {
		slog.Error("unable to read config", "error", err)
		os.Exit(1)
	}
	if err = json.Unmarshal(fileContent, &c); err != nil {
		slog.Error("unable to unmarshal config", "error", err)
		os.Exit(1)
	}
}

func NewConfig() *Config {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
	config := newDefaultConfig()
	config.readJsonConfig()
	if config.Bot.Identity.Token == "" {
		slog.Error("missing token", "error", "token is empty")
		os.Exit(1)
	}
	return config
}
