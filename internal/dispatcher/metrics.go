package dispatcher

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	TotalDispatchMessages = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "snorp_websocket_total_received_dispatch_messages",
		Help: "The total number of received websocket messages",
	},
		[]string{"action"},
	)
)

func init() {
	prometheus.MustRegister(TotalDispatchMessages)
}
