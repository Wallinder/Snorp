package server

import (
	"log/slog"
	"net/http"
)

func defaults(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("request", "method", r.Method, "uri", r.URL.Path)
		TotalServerHttpRequests.WithLabelValues(r.Method, r.URL.Path).Inc()
		next.ServeHTTP(w, r)
	})
}
