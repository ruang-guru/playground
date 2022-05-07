package middleware

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ruang-guru/playground/backend/performance-testing/server/metrics"
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
			metrics.IncrIncomingHTTPRequest(status, respCode)
		}
	}
}
