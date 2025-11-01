package metrics

import (
	"log"
	"net/http"
	"snorp/internal/state"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func MetricHandler(session *state.SessionState) *http.ServeMux {
	router := http.DefaultServeMux

	log.Printf("Metrics available at :%d%s\n", session.MetricPort, session.MetricUri)
	session.MetricRegistry = prometheus.NewRegistry()

	session.Metrics = NewMetrics(session.MetricRegistry)

	handlerOptions := promhttp.HandlerOpts{
		Registry: session.MetricRegistry,
	}
	router.Handle(session.MetricUri, promhttp.HandlerFor(session.MetricRegistry, handlerOptions))

	return router
}

func NewMetrics(reg prometheus.Registerer) state.Metrics {
	var metrics state.Metrics

	metrics.TotalReceivedMessages = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "websocket_total_received_messages",
		Help: "The total number of received websocket messages",
	})
	reg.MustRegister(metrics.TotalReceivedMessages)

	metrics.TotalDisconnects = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "websocket_total_disconnects",
		Help: "The total number of websocket disconnections",
	})
	reg.MustRegister(metrics.TotalDisconnects)

	return metrics
}
