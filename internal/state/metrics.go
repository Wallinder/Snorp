package state

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	TotalClientHttpRequests = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "snorp_http_total_client_requests",
		Help: "The total number of client requests",
	},
		[]string{"method", "path"},
	)
)

func init() {
	prometheus.MustRegister(TotalClientHttpRequests)
}
