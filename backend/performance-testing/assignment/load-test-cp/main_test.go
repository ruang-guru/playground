package main

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ruang-guru/playground/backend/performance-testing/server/api"
	"github.com/ruang-guru/playground/backend/performance-testing/server/handler"
	"github.com/ruang-guru/playground/backend/performance-testing/server/repository"

	"net/http"
	"net/http/httptest"
)

var _ = Describe("load test with vegeta library", func() {
	When("using attack command", func() {
		It("can get all movies", func() {
			w := setupRouter()
			ts := httptest.NewServer(w)
			defer ts.Close()
			sendFakeData(ts.URL)
			metrics := getMoviesTest(ts.URL)
			Expect(int(metrics.Requests)).To(Equal(20))
			Expect(metrics.StatusCodes["200"]).To(Equal(20))
		})

		It("can add new movie", func() {
			w := setupRouter()
			ts := httptest.NewServer(w)
			defer ts.Close()
			url := ts.URL
			metrics := addMovieTest(url)
			Expect(int(metrics.Requests)).To(Equal(10))
			Expect(metrics.StatusCodes["200"]).To(Equal(10))
		})
		It("can get movie by id", func() {
			w := setupRouter()
			ts := httptest.NewServer(w)
			defer ts.Close()
			sendFakeData(ts.URL)
			metrics := getMovieTest(ts.URL)
			Expect(int(metrics.Requests)).To(Equal(25))
			Expect(metrics.StatusCodes["200"]).To(Equal(25))
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

func sendFakeData(url string) {
	for i := 0; i < 40; i++ {
		payload, err := json.Marshal(Movie{
			Episode: i,
			Name:    "baru",
		})
		if err != nil {
			log.Fatal(err)
		}
		req, err := http.NewRequest("POST", url+"/movie", strings.NewReader(string(payload)))
		if err != nil {
			log.Fatal(err)
		}
		client := &http.Client{}
		client.Do(req)
	}
}
