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
	It("can have increase request rate", func() {
		w := setupRouter()
		ts := httptest.NewServer(w)
		defer ts.Close()
		metrics := increaseTest(ts.URL)
		Expect(int(metrics.Requests)).To(Equal(16))
		Expect(metrics.StatusCodes["200"]).To(Equal(16))
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
