package main

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Chaining Middleware", func() {
	Describe("hit endpoint / with GET request method and Header role is ADMIN", func() {
		It("should return OK", func() {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/", nil)
			r.Header.Set("role", "ADMIN")
			handler := RequestMethodGetMiddleware(AdminMiddleware(TimeHandler()))
			handler.ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusOK))
		})
	})
	Describe("hit endpoint / with POST request method and Header role is ADMIN", func() {
		It("should return method not allowed", func() {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("POST", "/", nil)
			r.Header.Set("role", "ADMIN")
			handler := RequestMethodGetMiddleware(AdminMiddleware(TimeHandler()))
			handler.ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusMethodNotAllowed))
		})
	})

	Describe("hit endpoint / with POST request method and no role", func() {
		It("should return method not allowed", func() {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("POST", "/", nil)
			handler := RequestMethodGetMiddleware(AdminMiddleware(TimeHandler()))
			handler.ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusMethodNotAllowed))
		})
	})

	Describe("hit endpoint / with GET request method and Header role is not ADMIN", func() {
		It("should return unauthorized", func() {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/", nil)
			r.Header.Set("role", "USER")
			handler := RequestMethodGetMiddleware(AdminMiddleware(TimeHandler()))
			handler.ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusUnauthorized))
		})
	})

	Describe("hit endpoint / with GET request method and no role", func() {
		It("should return unauthorized", func() {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/", nil)
			handler := RequestMethodGetMiddleware(AdminMiddleware(TimeHandler()))
			handler.ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusUnauthorized))
		})
	})
})
