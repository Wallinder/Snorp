package metrics

import (
	"log"
	"net/http"
	"snorp/internal/state"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func MetricHandler(session *state.SessionState) *http.ServeMux {
	router := http.DefaultServeMux

	NewMetrics(session)

	log.Printf("Metrics available at :%d%s\n", session.MetricPort, session.MetricUri)
	router.Handle(session.MetricUri, promhttp.Handler())

	return router
}

func NewMetrics(session *state.SessionState) {
	session.Metrics = &state.Metrics{
		TotalReceivedMessages: promauto.NewCounter(prometheus.CounterOpts{
			Name: "websocket_total_received_messages",
			Help: "The total number of received websocket messages",
		}),

		TotalDisconnects: promauto.NewCounter(prometheus.CounterOpts{
			Name: "websocket_total_disconnects",
			Help: "The total number of websocket disconnections",
		}),

		ActiveDispatchMessages: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "websocket_active_dispatch_messages",
			Help: "The number of active dispatch messages",
		}),
	}
}
