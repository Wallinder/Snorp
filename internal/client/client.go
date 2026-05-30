package client

import (
	"net/http"
	"time"
)

func NewHttpClient() *http.Client {
	return &http.Client{
		CheckRedirect: nil,
		Timeout:       5 * time.Second,
	}
}
