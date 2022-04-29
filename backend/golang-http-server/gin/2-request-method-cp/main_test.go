package main

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Gin Request Method", func() {
	Describe("hit endpoint /get with GET request method", func() {
		It("should return message GET", func() {
			route := GetGinRoute()
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/get", nil)
			route.ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(w.Body.String()).To(Equal("{\"message\":\"GET\"}"))
		})
	})
	Describe("hit endpoint /post with POST request method", func() {
		It("should return message POST", func() {
			route := GetGinRoute()
			w := httptest.NewRecorder()
			r := httptest.NewRequest("POST", "/post", nil)
			route.ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(w.Body.String()).To(Equal("{\"message\":\"POST\"}"))
		})
	})
	Describe("hit endpoint /put with PUT request method", func() {
		It("should return message PUT", func() {
			route := GetGinRoute()
			w := httptest.NewRecorder()
			r := httptest.NewRequest("PUT", "/put", nil)
			route.ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(w.Body.String()).To(Equal("{\"message\":\"PUT\"}"))
		})
	})
	Describe("hit endpoint /delete with DELETE request method", func() {
		It("should return message DELETE", func() {
			route := GetGinRoute()
			w := httptest.NewRecorder()
			r := httptest.NewRequest("DELETE", "/delete", nil)
			route.ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(w.Body.String()).To(Equal("{\"message\":\"DELETE\"}"))
		})
	})
	Describe("hit endpoint /patch with PATCH request method", func() {
		It("should return message PATCH", func() {
			route := GetGinRoute()
			w := httptest.NewRecorder()
			r := httptest.NewRequest("PATCH", "/patch", nil)
			route.ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(w.Body.String()).To(Equal("{\"message\":\"PATCH\"}"))
		})
	})
	Describe("hit endpoint /head with HEAD request method", func() {
		It("should return message HEAD", func() {
			route := GetGinRoute()
			w := httptest.NewRecorder()
			r := httptest.NewRequest("HEAD", "/head", nil)
			route.ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(w.Body.String()).To(Equal("{\"message\":\"HEAD\"}"))
		})
	})
	Describe("hit endpoint /options with OPTIONS request method", func() {
		It("should return message OPTIONS", func() {
			route := GetGinRoute()
			w := httptest.NewRecorder()
			r := httptest.NewRequest("OPTIONS", "/options", nil)
			route.ServeHTTP(w, r)
			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(w.Body.String()).To(Equal("{\"message\":\"OPTIONS\"}"))
		})
	})

})
