package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Routing HandleFunc", func() {
	Describe("hit endpoint /time with GET request method", func() {
		It("should return current day and date", func() {
			t := time.Now()
			expected := fmt.Sprintf("%v, %v %v %v", t.Weekday(), t.Day(), t.Month(), t.Year())

			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/time", nil)
			handler := http.HandlerFunc(TimeHandler)
			handler.ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(w.Body.String()).To(Equal(expected))
		})
	})

	Describe("hit endpoint /hello with GET request method", func() {
		It("should return Hello there", func() {

			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/hello", nil)
			handler := http.HandlerFunc(SayHelloHandler)
			handler.ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(w.Body.String()).To(Equal("Hello there"))
		})
	})

	Describe("hit endpoint /hello with GET request method and query param name", func() {
		It("should return Hello, name!", func() {

			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/hello?name=Roger", nil)
			handler := http.HandlerFunc(SayHelloHandler)
			handler.ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(w.Body.String()).To(Equal("Hello, Roger!"))
		})
	})
})
