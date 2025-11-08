package metrics

import (
	"context"
	"log"
	"net/http"
	"snorp/internal/state"
	"time"

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
		TotalMessages: promauto.NewCounterVec(prometheus.CounterOpts{
			Name: "websocket_total_received_messages",
			Help: "The total number of received websocket messages",
		},
			[]string{"opcode"},
		),

		TotalDispatchMessages: promauto.NewCounterVec(prometheus.CounterOpts{
			Name: "websocket_total_received_dispatch_messages",
			Help: "The total number of received websocket messages",
		},
			[]string{"action"},
		),

		TotalHttpRequests: promauto.NewCounterVec(prometheus.CounterOpts{
			Name: "http_total_client_requests",
			Help: "The total number of client requests",
		},
			[]string{"method", "path"},
		),

		TotalDisconnects: promauto.NewCounter(prometheus.CounterOpts{
			Name: "websocket_total_disconnects",
			Help: "The total number of websocket disconnections",
		}),

		AccumulatedMessages: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "websocket_accumelated_messages",
			Help: "The number of websocket messages within a timeperiod",
		}),
	}
}

func MessageMonitor(ctx context.Context, session *state.SessionState, messages chan []byte) {
	for {
		select {
		case <-ctx.Done():
			session.Metrics.AccumulatedMessages.Set(0)
			return
		case <-messages:
			session.Metrics.AccumulatedMessages.Inc()
			time.Sleep(3 * time.Second)
			session.Metrics.AccumulatedMessages.Dec()
		}
	}
}
