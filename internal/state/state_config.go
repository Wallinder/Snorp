package state

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"snorp/internal/models"
)

type Config struct {
	Bot DiscordBot `json:"discord_bot"`
}

type DiscordBot struct {
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
	Presence       models.Presence    `json:"presence"`
	Shards         []int              `json:"shards,omitempty"`
}

type IdentityProperties struct {
	Os      string `json:"os"`
	Browser string `json:"browser"`
	Device  string `json:"device"`
}

func newDefaultConfig() *Config {
	return &Config{
		Bot: DiscordBot{
			Permissions: 2031514586918128,
			Api:         "https://discord.com/api",
			ApiVersion:  "10",
			Identity: Identity{
				Properties: IdentityProperties{
					Os:      "Linux",
					Browser: "https://github.com/Wallinder/Snorp",
					Device:  "Walle-Lab",
				},
				Presence: models.Presence{
					Since: 0,
					Activities: []models.Activity{
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
		LogAndExit("unable to read config", err, 1)
	}
	if err = json.Unmarshal(fileContent, &c); err != nil {
		LogAndExit("unable to unmarshal config", err, 1)
	}
}

func NewConfig() *Config {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
	config := newDefaultConfig()
	config.readJsonConfig()
	if config.Bot.Identity.Token == "" {
		LogAndExit("no token was provided in config", fmt.Errorf("missing token"), 1)
	}
	return config
}
