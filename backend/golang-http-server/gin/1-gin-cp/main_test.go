package main

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Gin", func() {
	Describe("hit endpoint /hello with GET request method", func() {
		It("should return message hello", func() {
			route := GetGinRoute()
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/hello", nil)
			route.ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(w.Body.String()).To(Equal("{\"message\":\"hello\"}"))
		})
	})
	Describe("hit endpoint /world with GET request method", func() {
		It("should return message world", func() {
			route := GetGinRoute()
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/world", nil)
			route.ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(w.Body.String()).To(Equal("{\"message\":\"world\"}"))
		})
	})
})
