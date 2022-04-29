package main

import (
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ruang-guru/playground/backend/performance-testing/server/api"
	"github.com/ruang-guru/playground/backend/performance-testing/server/handler"
	"github.com/ruang-guru/playground/backend/performance-testing/server/repository"

	"net/http/httptest"
)

var _ = Describe("performance test with vegeta library", func() {
	It("can request like spike test", func() {
		w := setupRouter()
		ts := httptest.NewServer(w)
		defer ts.Close()
		metrics := spikeTest(ts.URL)
		Expect(int(metrics.Requests)).To(Equal(50))
		Expect(metrics.StatusCodes["200"]).To(Equal(50))
	})
})

func setupRouter() *gin.Engine {
	engine := gin.New()
	repo := repository.NewRepo()
	svc := handler.New(repo)
	serviceAPI := api.New(engine, svc)
	serviceAPI.InitAPI()
	return engine
}
