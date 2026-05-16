package discord

import (
	"fmt"
	"log/slog"
	"net/http"
)

type DiscordTransport struct {
	Token     string
	UserAgent string
	next      http.RoundTripper
}

func NewDiscordTransport(discordToken string, userAgent string) *DiscordTransport {
	return &DiscordTransport{
		next:      http.DefaultTransport,
		Token:     discordToken,
		UserAgent: userAgent,
	}
}

func (dt *DiscordTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req = req.Clone(req.Context())
	slog.Info("client request", "method", req.Method, "url", req.URL.Path)

	TotalClientHttpRequests.WithLabelValues(req.Method, req.URL.Path).Inc()

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", dt.UserAgent)
	req.Header.Set("Authorization", fmt.Sprintf("Bot %s", dt.Token))

	return dt.next.RoundTrip(req)
}
