package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/ruang-guru/playground/backend/performance-testing/server/api"
	"github.com/ruang-guru/playground/backend/performance-testing/server/handler"
	"github.com/ruang-guru/playground/backend/performance-testing/server/metrics"
	"github.com/ruang-guru/playground/backend/performance-testing/server/middleware"
	"github.com/ruang-guru/playground/backend/performance-testing/server/repository"
)

func main() {

	err := prometheus.Register(metrics.HTTPRequestCounter)
	if err != nil {
		panic("error registering prometheus metric")
	}
	engine := gin.New()
	engine.Use(middleware.MetricsMiddleware())

	repo := repository.NewRepo()

	svc := handler.New(repo)
	serviceAPI := api.New(engine, svc)
	serviceAPI.InitAPI()
	serviceAPI.InitPromeHandler()

	log.Fatal(engine.Run(":8090"))
}
