package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Bot DiscordBot `json:"discordBot"`
}

type DiscordBot struct {
	SuperuserID string   `json:"superuser_id"`
	Permissions int64    `json:"permissions"`
	Gateway     string   `json:"gateway"`
	Api         string   `json:"api"`
	ApiVersion  string   `json:"api_version"`
	Identity    Identity `json:"identity"`
	Mute        Mute     `json:"mute"`
}

type Identity struct {
	Token          string             `json:"token"`
	Compress       bool               `json:"compress"`
	LargeThreshold int                `json:"large_threshold"`
	Intents        int64              `json:"intents"`
	Properties     IdentityProperties `json:"properties"`
}

type IdentityProperties struct {
	Os      string `json:"os"`
	Browser string `json:"browser"`
	Device  string `json:"device"`
}

type Mute struct {
	Users []string `json:"users"`
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
