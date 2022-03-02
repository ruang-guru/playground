package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/ruangguru/playground/performance-testing/api"
	"github.com/ruangguru/playground/performance-testing/handlers"
	"github.com/ruangguru/playground/performance-testing/metrics"
	"github.com/ruangguru/playground/performance-testing/middleware"
	"github.com/ruangguru/playground/performance-testing/repository"
)

func main() {

	prometheus.Register(metrics.HttpRequestCounter)
	engine := gin.New()
	engine.Use(middleware.MetricsMiddleware())

	repo := repository.NewRepo()

	svc := handlers.New(repo)
	serviceApi := api.New(engine, svc)
	serviceApi.InitAPI()
	serviceApi.InitPromeHandler()

	engine.Run(":8090")
}
