package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type StaticConfig struct {
	Bot DiscordBot `yaml:"discordBot"`
}

type DiscordBot struct {
	Token       string   `yaml:"token"`
	Permissions int64    `yaml:"permissions"`
	Gateway     string   `yaml:"gateway"`
	Api         string   `yaml:"api"`
	Identity    Identity `yaml:"identity"`
}

type Identity struct {
	Compress       bool               `yaml:"compress"`
	LargeThreshold int                `yaml:"largethreshold"`
	Intents        int64              `yaml:"intents"`
	Properties     IdentityProperties `yaml:"properties"`
}

type IdentityProperties struct {
	Os      string `yaml:"os"`
	Browser string `yaml:"browser"`
	Device  string `yaml:"device"`
}

func Settings() StaticConfig {
	fileContent, err := os.ReadFile("../../config/config.yaml")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	var config StaticConfig
	err = yaml.Unmarshal([]byte(fileContent), &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return config
}
