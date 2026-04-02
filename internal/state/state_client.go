package state

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"
)

func newHttpClient(discordToken string) *http.Client {
	return &http.Client{
		CheckRedirect: nil,
		Timeout:       5 * time.Second,
		Transport:     newTransport(discordToken),
	}
}

func newTransport(discordToken string) *discordRoundTripper {
	return &discordRoundTripper{
		next:  http.DefaultTransport,
		token: discordToken,
	}
}

type discordRoundTripper struct {
	next  http.RoundTripper
	token string
}

func (drt *discordRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	slog.Info("discord request", "method", req.Method, "url", req.URL.Path)

	TotalClientHttpRequests.WithLabelValues(req.Method, req.URL.Path).Inc()

	req.Header = map[string][]string{
		"Content-Type":  {"application/json"},
		"User-Agent":    {"DiscordBot (https://github.com/Wallinder/Snorp)"},
		"Authorization": {fmt.Sprintf("Bot %s", drt.token)},
	}

	resp, err := drt.next.RoundTrip(req)
	return resp, err
}

func (s *SessionState) NewDiscordRequest(method string, uri string, body io.Reader) (*http.Response, error) {
	url := s.Config.Bot.Api + "/v" + s.Config.Bot.ApiVersion + uri

	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	response, err := s.Client.Do(request)
	return response, err
}
