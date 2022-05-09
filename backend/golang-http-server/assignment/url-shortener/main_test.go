package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/url-shortener/entity"
	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/url-shortener/handlers"
	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/url-shortener/repository"
	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/url-shortener/router"
)

var _ = Describe("URL Shortener", func() {
	When("Get non existing shorturl", func() {
		It("should return not found", func() {
			repo := repository.NewMapRepository()
			h := handlers.NewURLHandler(&repo)
			r := router.SetupRouter(h)
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/nonexist", nil)
			r.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusNotFound))
		})
	})
	When("Shorten longurl with random shorturl", func() {
		It("should return no error", func() {
			var b bytes.Buffer
			json.NewEncoder(&b).Encode(entity.URL{LongURL: "https://pawgrammers.com"})

			repo := repository.NewMapRepository()
			h := handlers.NewURLHandler(&repo)
			r := router.SetupRouter(h)
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/", &b)
			r.ServeHTTP(w, req)

			var res struct {
				Data entity.URL
			}
			json.Unmarshal(w.Body.Bytes(), &res)
			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(res.Data.LongURL).To(Equal("https://pawgrammers.com"))
		})
	})
	When("Shorten longurl with custom shorturl", func() {
		It("should return no error", func() {
			var b bytes.Buffer
			json.NewEncoder(&b).Encode(entity.URL{LongURL: "https://pawgrammers.com", ShortURL: "pawgrammers"})

			repo := repository.NewMapRepository()
			h := handlers.NewURLHandler(&repo)
			r := router.SetupRouter(h)
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/custom", &b)
			r.ServeHTTP(w, req)

			var res struct {
				Data entity.URL
			}
			json.Unmarshal(w.Body.Bytes(), &res)
			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(res.Data.LongURL).To(Equal("https://pawgrammers.com"))
			Expect(res.Data.ShortURL).To(Equal("pawgrammers"))
		})
	})
	When("Shorten longurl with existing custom shorturl", func() {
		It("should return error", func() {
			var b bytes.Buffer
			json.NewEncoder(&b).Encode(entity.URL{LongURL: "https://pawgrammers.com", ShortURL: "pawgrammers"})

			repo := repository.NewMapRepository()
			h := handlers.NewURLHandler(&repo)
			r := router.SetupRouter(h)
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/custom", &b)
			r.ServeHTTP(w, req)

			w = httptest.NewRecorder()
			req = httptest.NewRequest("POST", "/custom", &b)
			r.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusBadRequest))
		})
	})
	When("Get existing short url", func() {
		It("should redirect to long url", func() {
			var b bytes.Buffer
			json.NewEncoder(&b).Encode(entity.URL{LongURL: "https://pawgrammers.com", ShortURL: "pawgrammers"})

			repo := repository.NewMapRepository()
			h := handlers.NewURLHandler(&repo)
			r := router.SetupRouter(h)
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/custom", &b)
			r.ServeHTTP(w, req)

			w = httptest.NewRecorder()
			req = httptest.NewRequest("GET", "/pawgrammers", &b)
			r.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusFound))
			Expect(w.HeaderMap.Get("Location")).To(Equal("https://pawgrammers.com"))
		})
	})
})
