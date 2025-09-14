package api

import (
	"fmt"
	"net/http"
)

func CreateVoiceChannel(api, token string) error {
	req, err := http.NewRequest("GET", api, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bot %s", token))
	req.Header.Add("User-Agent", "")
	//resp, err := client.Do(req)
	return nil
}
