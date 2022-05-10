package main

import (
	"errors"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Error Handling", func() {
	Describe("hit endpoint / with GET request method with empty data and err nil", func() {
		It("should return Status Not Found", func() {
			data = []string{}
			err = nil
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/", nil)
			handler.ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusNotFound))
		})
	})
	Describe("hit endpoint / with GET request method with non empty data and err nil", func() {
		It("should return OK", func() {
			data = []string{"data1"}
			err = nil
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/", nil)
			handler.ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusOK))
		})
	})
	Describe("hit endpoint / with GET request method with non empty data and err not nil", func() {
		It("should return Internal Server Error", func() {
			data = []string{"data1"}
			err = errors.New("error internal server error")
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/", nil)
			handler.ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusInternalServerError))
		})
	})
})
