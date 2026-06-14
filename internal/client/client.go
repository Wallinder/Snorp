package client

import (
	"log/slog"
	"net/http"
	"time"
)

func NewHttpClient() *http.Client {
	return &http.Client{
		CheckRedirect: nil,
		Timeout:       5 * time.Second,
		Transport: &CustomTransport{
			Next: http.DefaultTransport,
		},
	}
}

type CustomTransport struct {
	Next http.RoundTripper
}

func (ct *CustomTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	slog.Info("client request", "method", req.Method, "url", req.URL.Path)
	return ct.Next.RoundTrip(req)
}
