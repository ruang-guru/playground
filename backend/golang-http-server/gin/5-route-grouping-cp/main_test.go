package main

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Gin Route Grouping", func() {
	Describe("hit endpoint /movie/list with GET request method", func() {
		It("should return OK and list of movie", func() {
			route := GetRouter()
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/movie/list", nil)
			route.ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusOK))
		})
	})
	Describe("hit endpoint /movie/get with GET request method and exists movie", func() {
		It("should return OK and movie", func() {
			route := GetRouter()
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/movie/get/1", nil)
			route.ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(w.Body.String()).To(Equal("{\"data\":{\"Title\":\"Spiderman\"}}"))
		})
	})
	Describe("hit endpoint /movie/get with GET request method and not exists movie", func() {
		It("should return Not Found", func() {
			route := GetRouter()
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/movie/get/10", nil)
			route.ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusNotFound))
		})
	})
})
