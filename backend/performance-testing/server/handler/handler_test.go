package handler_test

import (
	"encoding/json"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ruang-guru/playground/backend/performance-testing/server/api"
	"github.com/ruang-guru/playground/backend/performance-testing/server/handler"
	"github.com/ruang-guru/playground/backend/performance-testing/server/repository"

	"log"
	"net/http"
	"net/http/httptest"
)

type Movie struct {
	ID      int    `json:"id"`
	Episode int    `json:"episode"`
	Name    string `json:"name"`
}

var router *gin.Engine
var _ = Describe("handler", func() {
	When("the handler implemented", func() {
		It("can add new movie with method POST", func() {
			var wg sync.WaitGroup
			var mtx sync.Mutex
			w := httptest.NewRecorder()
			for i := 0; i < 1000; i++ {
				wg.Add(1)
				go func(i int) {
					defer wg.Done()
					payload, err := json.Marshal(Movie{
						Episode: i,
						Name:    "baru",
					})
					if err != nil {
						log.Fatal(err)
					}
					req, err := http.NewRequest("POST", "/movie", strings.NewReader(string(payload)))
					if err != nil {
						log.Fatal(err)
					}
					mtx.Lock()
					router.ServeHTTP(w, req)
					mtx.Unlock()
				}(i)
			}
			wg.Wait()
			mtx.Lock()
			Expect(w.Code).To(Equal(200))
			mtx.Unlock()
		})

		It("can get all movies", func() {
			w := httptest.NewRecorder()
			gin.CreateTestContext(w)
			req, err := http.NewRequest("GET", "/movies", nil)
			if err != nil {
				log.Fatal(err)
			}
			router.ServeHTTP(w, req)
			var movies []Movie
			err = json.NewDecoder(w.Body).Decode(&movies)
			if err != nil {
				log.Fatal(err)
			}
			Expect(w.Code).To(Equal(200))
			Expect(len(movies)).To(Equal(1000))
		})

		It("can get a movie", func() {
			w := httptest.NewRecorder()

			req, err := http.NewRequest("GET", "/movie/1", nil)
			if err != nil {
				log.Fatal(err)
			}
			router.ServeHTTP(w, req)

			var movie Movie
			err = json.NewDecoder(w.Body).Decode(&movie)

			if err != nil {
				log.Fatal(err)
			}
			Expect(w.Code).To(Equal(200))
			Expect(movie.ID).To(Equal(1))
		})

		It("can test GET method", func() {
			w := httptest.NewRecorder()

			req, err := http.NewRequest("GET", "/", nil)
			if err != nil {
				log.Fatal(err)
			}
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(200))
		})

		It("can test POST method", func() {
			w := httptest.NewRecorder()
			r := strings.NewReader("hi")
			req, err := http.NewRequest("POST", "/", r)
			if err != nil {
				log.Fatal(err)
			}
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(200))
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

func init() {
	router = setupRouter()
}
