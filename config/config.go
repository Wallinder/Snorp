package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Bot        DiscordBot `json:"discordbot"`
	Postgresql Postgresql `json:"postgresql"`
	SVV        SVV        `json:"svv"`
}

type SVV struct {
	ApiKey string `json:"api_key"`
}

type Postgresql struct {
	ConnectionString string `json:"connection_string"`
	Gorm             Gorm   `json:"gorm"`
}

type Gorm struct {
	SingularTable   bool `json:"singular_table"`
	MaxIdleConns    int  `json:"max_idle_conns"`
	MaxOpenConns    int  `json:"max_open_conns"`
	ConnMaxLifetime int  `json:"conn_max_lifetime"`
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
	Presence       Presence           `json:"presence"`
	Shards         []int              `json:"shards,omitempty"`
}

type IdentityProperties struct {
	Os      string `json:"os"`
	Browser string `json:"browser"`
	Device  string `json:"device"`
}

type Presence struct {
	Since      int        `json:"since"`
	Activities []Activity `json:"activities"`
	Status     string     `json:"status"`
	AFK        bool       `json:"afk"`
}

type Activity struct {
	Name string `json:"name"`
	Type int    `json:"type"`
}

func Settings() Config {
	fileContent, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	var config Config

	err = json.Unmarshal([]byte(fileContent), &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	if config.Bot.Identity.Token == "" {
		log.Fatal("Missing token..")
	}

	return config
}
