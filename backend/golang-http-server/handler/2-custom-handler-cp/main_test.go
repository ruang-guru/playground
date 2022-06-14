package main

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Handler", func() {
	Describe("hit endpoint / with GET request method", func() {
		It("should return current day and time", func() {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/", nil)
			handler := QuotesHandler{}
			handler.ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(quotes).Should(ContainElement(w.Body.String()))
		})
	})
	Describe("hit endpoint / with POST request method", func() {
		It("should return current day and time", func() {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("POST", "/", nil)
			handler := QuotesHandler{}
			handler.ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(quotes).Should(ContainElement(w.Body.String()))
		})
	})
	Describe("hit endpoint /quote with GET request method", func() {
		It("should return current day and time", func() {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/quote", nil)
			handler := QuotesHandler{}
			handler.ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(quotes).Should(ContainElement(w.Body.String()))
		})
	})
})
