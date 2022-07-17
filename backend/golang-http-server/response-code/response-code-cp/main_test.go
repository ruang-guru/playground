package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Handler", func() {
	Describe("hit endpoint /name with GET request method and non exists name", func() {
		It("should return status not found", func() {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/name?name=Sophia", nil)

			GetMux().ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusNotFound))
		})
	})
	Describe("hit endpoint /name with POST request method", func() {
		It("should return status method not allowed", func() {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("POST", "/name?name=Roger", nil)
			GetMux().ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusMethodNotAllowed))
		})
	})
	Describe("hit endpoint /name with GET request method and exists name", func() {
		It("should return status OK", func() {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/name?name=Roger", nil)
			fmt.Println(r.URL.Query())

			GetMux().ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusOK))
		})
	})
})
