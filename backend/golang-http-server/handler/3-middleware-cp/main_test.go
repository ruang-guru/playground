package main

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Middleware", func() {
	Describe("hit endpoint / with GET request method", func() {
		It("should return OK", func() {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/", nil)
			timeHandler := TimeHandler{}
			handler := RequestMethodGetMiddleware(timeHandler)
			handler.ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusOK))
		})
	})
	Describe("hit endpoint / with POST request method", func() {
		It("should return method not allowed", func() {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("POST", "/", nil)
			timeHandler := TimeHandler{}
			handler := RequestMethodGetMiddleware(timeHandler)
			handler.ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusMethodNotAllowed))
		})
	})
	Describe("hit endpoint / with PUT request method", func() {
		It("should return method not allowed", func() {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("PUT", "/", nil)
			timeHandler := TimeHandler{}
			handler := RequestMethodGetMiddleware(timeHandler)
			handler.ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusMethodNotAllowed))
		})
	})
	Describe("hit endpoint / with DELET request method", func() {
		It("should return method not allowed", func() {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("DELETE", "/", nil)
			timeHandler := TimeHandler{}
			handler := RequestMethodGetMiddleware(timeHandler)
			handler.ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusMethodNotAllowed))
		})
	})
})
