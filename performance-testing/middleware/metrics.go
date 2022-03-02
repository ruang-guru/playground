package middleware

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ruangguru/playground/performance-testing/metrics"
)

func MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		var status string
		respCode := strconv.Itoa(c.Writer.Status())
		if c.Writer.Status() >= 300 {
			status = "error"
		} else {
			status = "success"
		}

		if c.FullPath() != "/metrics" {
			metrics.IncrIncomingHttpRequest(status, respCode)
		}
	}
}
