package api

import (
	"fmt"
)

func GetHeaders(token string) map[string][]string {
	return map[string][]string{
		"Content-Type":  {"application/json"},
		"User-Agent":    {"DiscordBot (https://github.com/Wallinder/Snorp; latest)"},
		"Authorization": {fmt.Sprintf("Bot %s", token)},
	}
}
