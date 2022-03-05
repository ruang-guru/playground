package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/ruang-guru/playground/backend/performance-testing/api"
	"github.com/ruang-guru/playground/backend/performance-testing/handlers"
	"github.com/ruang-guru/playground/backend/performance-testing/metrics"
	"github.com/ruang-guru/playground/backend/performance-testing/middleware"
	"github.com/ruang-guru/playground/backend/performance-testing/repository"
)

func main() {

	err := prometheus.Register(metrics.HTTPRequestCounter)
	if err != nil {
		panic("error registering prometheus metric")
	}
	engine := gin.New()
	engine.Use(middleware.MetricsMiddleware())

	repo := repository.NewRepo()

	svc := handlers.New(repo)
	serviceAPI := api.New(engine, svc)
	serviceAPI.InitAPI()
	serviceAPI.InitPromeHandler()

	log.Fatal(engine.Run(":8090"))
}
