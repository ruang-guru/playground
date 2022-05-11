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

var _ = Describe("load test with vegeta library", func() {
	When("using attack command", func() {
		It("can use GET method", func() {
			w := setupRouter()
			ts := httptest.NewServer(w)
			defer ts.Close()
			metrics := vegetaGet(ts.URL)
			Expect(int(metrics.Requests)).To(Equal(20))
			Expect(metrics.StatusCodes["200"]).To(Equal(20))
		})

		It("can use POST method", func() {
			w := setupRouter()
			ts := httptest.NewServer(w)
			defer ts.Close()
			url := ts.URL
			metrics := vegetaPost(url)
			Expect(int(metrics.Requests)).To(Equal(30))
			Expect(metrics.StatusCodes["200"]).To(Equal(30))
		})
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
