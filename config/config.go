package config

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"snorp/internal/storage"
	"snorp/pkg/discord"
)

type Config struct {
	Bot      DiscordBot           `json:"discord_bot"`
	Postgres storage.PostgresOpts `json:"postgres"`
}

type DiscordBot struct {
	Permissions int64            `json:"permissions"`
	Gateway     string           `json:"gateway"`
	Api         string           `json:"api"`
	ApiVersion  string           `json:"api_version"`
	Identity    discord.Identity `json:"identity"`
}

func newDefaultConfig() *Config {
	return &Config{
		Bot: DiscordBot{
			Permissions: 2031514586918128,
			Api:         "https://discord.com/api",
			ApiVersion:  "10",
			Identity: discord.Identity{
				Properties: discord.IdentityProperties{
					Os:      "Linux",
					Browser: "https://github.com/Wallinder/Snorp",
					Device:  "Walle-Lab",
				},
				Presence: discord.Presence{
					Since: 0,
					Activities: []*discord.Activity{
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
		Postgres: storage.PostgresOpts{
			Enabled:           true,
			ConnectionString:  os.Getenv("PG_CONNECTION_STRING"),
			MaxConns:          20,
			MinConns:          5,
			MaxConnLifetime:   1800,
			MaxConnIdleTime:   900,
			HealthcheckPeriod: 60,
		},
	}
}

func (c *Config) readJsonConfig() error {
	fileContent, err := os.ReadFile("config.json")
	if err != nil {
		return err
	}
	return json.Unmarshal(fileContent, &c)
}

func newSlogHandler() *slog.JSONHandler {
	return slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.String(a.Key, a.Value.Time().Format("2006-01-02 15:04:05"))
			}
			return a
		},
	})
}

func NewConfig() (*Config, error) {
	slog.SetDefault(slog.New(newSlogHandler()))
	config := newDefaultConfig()
	if err := config.readJsonConfig(); err != nil {
		return nil, err
	}
	if config.Bot.Identity.Token == "" {
		return nil, fmt.Errorf("missing token")
	}
	return config, nil
}
