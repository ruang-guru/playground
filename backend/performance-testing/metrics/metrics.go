package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	unit      = "unit"
	status    = "status"
	respCode  = "response_code"
	queueName = "queue_name"

	httpCounterMetrics     = "http_request_count"
	httpCounterMetricsHelp = "Total count incoming request"
)

var (
	HttpRequestCounter *prometheus.CounterVec
)

func init() {
	HttpRequestCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: httpCounterMetrics,
		Help: httpCounterMetricsHelp,
	},
		[]string{status, respCode},
	)
}

func IncrIncomingHttpRequest(status string, respCode string) {
	HttpRequestCounter.WithLabelValues(status, respCode).Inc()
}
