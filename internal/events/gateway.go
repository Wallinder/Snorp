package event

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Metadata struct {
	Url               string  `json:"url"`
	Shards            int     `json:"shards"`
	SessionStartLimit Session `json:"session_start_limit"`
}

type Session struct {
	Total          int `json:"total"`
	Remaining      int `json:"remaining"`
	ResetAfter     int `json:"reset_after"`
	MaxConcurrency int `json:"max_concurrency"`
}

func GetGateway(gateway string, token string) Metadata {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

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

	var md Metadata
	err = json.Unmarshal(body, &md)
	if err != nil {
		log.Fatal(err)
	}

	return md
}
