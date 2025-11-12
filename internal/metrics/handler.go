package metrics

import (
	"log"
	"net/http"
	"snorp/internal/state"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Collector(session *state.SessionState) {
	NewMetrics(session)

	log.Printf("Metrics available at :%d%s\n", session.MetricPort, session.MetricUri)
	http.Handle(session.MetricUri, promhttp.Handler())

}

func NewMetrics(session *state.SessionState) {
	session.Metrics = &state.Metrics{
		TotalMessages: promauto.NewCounterVec(prometheus.CounterOpts{
			Name: "snorp_websocket_total_received_messages",
			Help: "The total number of received websocket messages",
		},
			[]string{"opcode"},
		),

		TotalDispatchMessages: promauto.NewCounterVec(prometheus.CounterOpts{
			Name: "snorp_websocket_total_received_dispatch_messages",
			Help: "The total number of received websocket messages",
		},
			[]string{"action"},
		),

		TotalHttpRequests: promauto.NewCounterVec(prometheus.CounterOpts{
			Name: "snorp_http_total_client_requests",
			Help: "The total number of client requests",
		},
			[]string{"method", "path"},
		),

		TotalDisconnects: promauto.NewCounter(prometheus.CounterOpts{
			Name: "snorp_websocket_total_disconnects",
			Help: "The total number of websocket disconnections",
		}),
	}
}
