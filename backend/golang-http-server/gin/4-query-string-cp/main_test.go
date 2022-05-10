package main

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Gin Path Parameter", func() {
	Describe("hit endpoint /movie with GET request method with genre and country query param", func() {
		It("should return output that contain genre and country", func() {
			route := GetRouter()
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/movie?genre=action&country=ID", nil)
			route.ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(w.Body.String()).To(Equal("Here result of query of movie with genre action in country ID"))
		})
	})
	Describe("hit endpoint /movie with GET request method country query param", func() {
		It("should return output that contain default genre and country", func() {
			route := GetRouter()
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/movie?country=ID", nil)
			route.ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(w.Body.String()).To(Equal("Here result of query of movie with genre general in country ID"))
		})
	})
	Describe("hit endpoint /movie with GET request method without query param", func() {
		It("should return output that contain default genre", func() {
			route := GetRouter()
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/movie", nil)
			route.ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(w.Body.String()).To(Equal("Here result of query of movie with genre general"))
		})
	})
	Describe("hit endpoint /movie with GET request method with genre query param", func() {
		It("should return output that contain default genre", func() {
			route := GetRouter()
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/movie?genre=horror", nil)
			route.ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(w.Body.String()).To(Equal("Here result of query of movie with genre horror"))
		})
	})
})
