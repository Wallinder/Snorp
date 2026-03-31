package event

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	TotalMessages = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "snorp_websocket_total_received_messages",
		Help: "The total number of received websocket messages",
	},
		[]string{"opcode"},
	)
)

func init() {
	prometheus.MustRegister(TotalMessages)
}
