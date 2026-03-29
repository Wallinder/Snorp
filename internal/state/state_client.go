package state

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"
)

func (session *SessionState) newHttpClient() {
	session.Client = &http.Client{
		CheckRedirect: nil,
		Timeout:       time.Duration(5 * time.Second),
		Transport:     newTransport(session),
	}
}

func newTransport(session *SessionState) *discordRoundTripper {
	return &discordRoundTripper{
		next:  http.DefaultTransport,
		state: session,
	}
}

type discordRoundTripper struct {
	next  http.RoundTripper
	state *SessionState
}

func (drt *discordRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	slog.Info("discord request", "method", req.Method, "url", req.URL.Path)

	drt.state.Metrics.TotalHttpRequests.WithLabelValues(req.Method, req.URL.Path).Inc()
	req.Header = map[string][]string{
		"Content-Type":  {"application/json"},
		"User-Agent":    {"DiscordBot (https://github.com/Wallinder/Snorp)"},
		"Authorization": {fmt.Sprintf("Bot %s", drt.state.Config.Bot.Identity.Token)},
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
