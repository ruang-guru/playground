package main

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Gin Path Parameter", func() {
	Describe("hit endpoint /profile with GET request method and not exists data", func() {
		It("should return not found error", func() {
			route := GetRouter()
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/profile/4", nil)
			route.ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusNotFound))
			Expect(w.Body.String()).To(Equal("data not found"))
		})
	})
	Describe("hit endpoint /profile with GET request method and exists data", func() {
		It("should return OK", func() {
			route := GetRouter()
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/profile/1", nil)
			route.ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(w.Body.String()).To(Equal("Name: Roger, Country: United States, Age: 70"))
		})
	})
	Describe("hit endpoint /profile with GET request method and exists data", func() {
		It("should return OK", func() {
			route := GetRouter()
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/profile/2", nil)
			route.ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(w.Body.String()).To(Equal("Name: Tony, Country: United States, Age: 40"))
		})
	})
})
