package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/ruang-guru/playground/backend/performance-testing/api"
	"github.com/ruang-guru/playground/backend/performance-testing/handlers"
	"github.com/ruang-guru/playground/backend/performance-testing/metrics"
	"github.com/ruang-guru/playground/backend/performance-testing/middleware"
	"github.com/ruang-guru/playground/backend/performance-testing/repository"
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
