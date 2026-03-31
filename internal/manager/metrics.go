package manager

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	TotalDisconnects = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "snorp_websocket_total_disconnects",
		Help: "The total number of websocket disconnections",
	})
)

func init() {
	prometheus.MustRegister(TotalDisconnects)
}
