package discord

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	TotalClientHttpRequests = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "snorp_discord_http_total_client_requests",
		Help: "The total number of client requests",
	},
		[]string{"method", "path"},
	)
	TotalDisconnects = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "snorp_websocket_total_disconnects",
		Help: "The total number of websocket disconnections",
	})
	TotalMessages = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "snorp_websocket_total_received_messages",
		Help: "The total number of received websocket messages",
	},
		[]string{"opcode"},
	)
	TotalDispatchMessages = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "snorp_websocket_total_received_dispatch_messages",
		Help: "The total number of received websocket messages",
	},
		[]string{"action"},
	)
)

func init() {
	prometheus.MustRegister(TotalClientHttpRequests)
	prometheus.MustRegister(TotalDisconnects)
	prometheus.MustRegister(TotalMessages)
	prometheus.MustRegister(TotalDispatchMessages)
}
