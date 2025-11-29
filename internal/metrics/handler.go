package metrics

import (
	"fmt"
	"log"
	"net/http"
	"snorp/internal/state"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Collector(metrics *state.Metrics) {
	NewMetrics(metrics)

	log.Printf("Metrics available at :%d%s\n", metrics.Port, metrics.Uri)
	http.Handle(metrics.Uri, promhttp.Handler())

	http.ListenAndServe(fmt.Sprintf(":%d", metrics.Port), nil)
}

func NewMetrics(metrics *state.Metrics) {
	metrics.TotalMessages = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "snorp_websocket_total_received_messages",
		Help: "The total number of received websocket messages",
	},
		[]string{"opcode"},
	)

	metrics.TotalDispatchMessages = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "snorp_websocket_total_received_dispatch_messages",
		Help: "The total number of received websocket messages",
	},
		[]string{"action"},
	)

	metrics.TotalHttpRequests = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "snorp_http_total_client_requests",
		Help: "The total number of client requests",
	},
		[]string{"method", "path"},
	)

	metrics.TotalDisconnects = promauto.NewCounter(prometheus.CounterOpts{
		Name: "snorp_websocket_total_disconnects",
		Help: "The total number of websocket disconnections",
	})
}
