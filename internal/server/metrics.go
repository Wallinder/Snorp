package server

import "github.com/prometheus/client_golang/prometheus"

var (
	TotalServerHttpRequests = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "snorp_http_total_server_requests",
		Help: "The total number of server requests",
	},
		[]string{"method", "path"},
	)
)

func init() {
	prometheus.MustRegister(TotalServerHttpRequests)
}
