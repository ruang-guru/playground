package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	status   = "status"
	respCode = "response_code"

	httpCounterMetrics     = "http_request_count"
	httpCounterMetricsHelp = "Total count incoming request"
)

var (
	HTTPRequestCounter *prometheus.CounterVec
)

func init() {
	HTTPRequestCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: httpCounterMetrics,
		Help: httpCounterMetricsHelp,
	},
		[]string{status, respCode},
	)
}

func IncrIncomingHTTPRequest(status string, respCode string) {
	HTTPRequestCounter.WithLabelValues(status, respCode).Inc()
}
