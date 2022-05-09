package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Handler", func() {
	Describe("hit endpoint /time with GET request method", func() {
		It("should return current day and time", func() {
			t := time.Now()
			expected := fmt.Sprintf("%v, %v %v %v", t.Weekday(), t.Day(), t.Month(), t.Year())

			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/time", nil)
			GetMux().ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(w.Body.String()).To(Equal(expected))
		})
	})

	Describe("hit endpoint /hello with GET request method", func() {
		It("should return Hello there", func() {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/hello", nil)
			GetMux().ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(w.Body.String()).To(Equal("Hello there"))
		})
	})
	Describe("hit endpoint /hello with GET request method and query param Roger", func() {
		It("should return Hello, Roger!", func() {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/hello?name=Roger", nil)
			GetMux().ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(w.Body.String()).To(Equal("Hello, Roger!"))
		})
	})

	Describe("hit endpoint /custom with GET request method", func() {
		It("should return not found", func() {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/custom", nil)
			GetMux().ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusNotFound))
		})
	})
})
