package config

import (
	"log"
	"os"
	"gopkg.in/yaml.v3"
)

type StaticConfig struct {
	Bot struct {
  		Token string `yaml:"token"`
  		Permissions int64 `yaml:"permissions"`
	}
	Url struct {
  		Gateway string `yaml:"gateway"`
  		Api string `yaml:"api"`
	}
	Identity struct {
  		Properties struct {
	    		Os  string `yaml:"os"`
	    		Browser string `yaml:"browser"`
	    		Device string `yaml:"device"`
		}
  		Compress bool `yaml:"compress"`
  		LargeThreshold int `yaml:"largethreshold"`
  		Intents int64 `yaml:"intents"`
	}
}

func Settings() StaticConfig {
	fileContent, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	var config StaticConfig
	_, err := yaml.Unmarshal([]byte(fileContent), &config)
	if err != nil {
		log.Fatalf("error: %v", err)
    }
	return config
}
