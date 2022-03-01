package handlers_test

import (
	"encoding/json"
	"strings"

	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ruangguru/playground/performance-testing/handlers"

	"log"
	"net/http"
	"net/http/httptest"
)

type Movie struct {
	Id      int    `json:"id"`
	Episode int    `json:"episode"`
	Name    string `json:"name"`
}

var _ = Describe("handlers", func() {
	When("the handlers implemented", func() {
		It("can add new movie with method POST", func() {
			ts := httptest.NewServer(NewHandler())
			defer ts.Close()

			payload, err := json.Marshal(Movie{
				Episode: 1,
				Name:    "baru",
			})

			if err != nil {
				log.Fatal(err)
			}

			res, err := http.Post(ts.URL+"/movie", "", strings.NewReader(string(payload)))
			if err != nil {
				log.Fatal(err)
			}
			Expect(res.StatusCode).To(Equal(200))
			Expect(res.StatusCode).To(Equal(200))
		})

		It("can get all movies", func() {
			ts := httptest.NewServer(NewHandler())
			defer ts.Close()

			res, err := http.Get(ts.URL + "/movies")
			if err != nil {
				log.Fatal(err)
			}

			var movies []Movie
			err = json.NewDecoder(res.Body).Decode(&movies)
			res.Body.Close()
			if err != nil {
				log.Fatal(err)
			}
			Expect(res.StatusCode).To(Equal(200))
			Expect(len(movies)).To(Equal(1))
		})

		It("can get a movie", func() {
			ts := httptest.NewServer(NewHandler())
			defer ts.Close()

			res, err := http.Get(ts.URL + "/movie/1")
			if err != nil {
				log.Fatal(err)
			}

			var movie Movie
			err = json.NewDecoder(res.Body).Decode(&movie)
			res.Body.Close()
			if err != nil {
				log.Fatal(err)
			}
			Expect(res.StatusCode).To(Equal(200))
			Expect(movie.Id).To(Equal(1))
		})

		It("can test GET method", func() {
			ts := httptest.NewServer(NewHandler())
			defer ts.Close()

			res, err := http.Get(ts.URL)
			if err != nil {
				log.Fatal(err)
			}

			if err != nil {
				log.Fatal(err)
			}

			Expect(res.StatusCode).To(Equal(200))
		})

		It("can test POST method", func() {
			ts := httptest.NewServer(NewHandler())
			defer ts.Close()

			res, err := http.Get(ts.URL)
			if err != nil {
				log.Fatal(err)
			}

			res.Body.Close()
			if err != nil {
				log.Fatal(err)
			}

			Expect(res.StatusCode).To(Equal(200))
		})
	})
})

func NewHandler() http.Handler {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/movies", handlers.GetMovies)
	router.HandleFunc("/movie", handlers.AddMovie)
	router.HandleFunc("/movie/{id}", handlers.GetMovie)
	router.HandleFunc("/", handlers.Home)
	return router
}
