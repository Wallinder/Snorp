package state

import (
	"net/http"
	"time"
)

func (s *SessionState) RunHttpServer(handler http.Handler) {
	s.Server.Handler = handler
	if err := s.Server.ListenAndServe(); err != nil {
		LogAndExit("http panic", err, 1)
	}
}

func newHttpServer() *http.Server {
	return &http.Server{
		Addr:              ":8080",
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       5 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}
}
