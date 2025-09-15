package api

import (
	"fmt"
	"io"
	"net/http"
)

func CreateRequest(method, url, token string, body io.Reader, client *http.Client) error {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}
	headers := map[string]string{
		"Content-Type":  "application/json",
		"User-Agent":    "DiscordBot (https://github.com/Wallinder/Snorp; latest)",
		"Authorization": fmt.Sprintf("Bot %s", token),
	}

	var reqHeaders http.Header

	for key, value := range headers {
		reqHeaders.Set(key, value)
	}

	_, err = client.Do(req)
	if err != nil {
		return err
	}

	return nil
}
