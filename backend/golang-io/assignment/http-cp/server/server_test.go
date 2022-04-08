package main_test

import (
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	server "github.com/ruang-guru/playground/backend/golang-io/assignment/http-cp/server"
)

var _ = Describe("Server", func() {
	Describe("/", func() {
		It("should return Hello, world!", func() {
			wr := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/", nil)
			server.Routes().ServeHTTP(wr, req)

			Expect(wr.Code).To(Equal(200))
			Expect(wr.Body.String()).To(Equal("Hello, world!"))
		})

		It("should return status OK", func() {
			wr := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/", nil)
			server.Routes().ServeHTTP(wr, req)

			Expect(wr.Code).To(Equal(200))
		})
	})
	Describe("/echo", func() {
		It("should return back the 'data' query param", func() {
			wr := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/echo?data=Hi!", nil)
			server.Routes().ServeHTTP(wr, req)

			Expect(wr.Code).To(Equal(200))
			Expect(wr.Body.String()).To(Equal("Hi!"))

			wr = httptest.NewRecorder()
			req = httptest.NewRequest("GET", "/echo?data=Happy", nil)
			server.Routes().ServeHTTP(wr, req)

			Expect(wr.Code).To(Equal(200))
			Expect(wr.Body.String()).To(Equal("Happy"))
		})
	})

	Describe("/add", func() {
		It("should return the sum of query param 'a' and 'b'", func() {
			wr := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/add?a=2&b=3", nil)
			server.Routes().ServeHTTP(wr, req)

			Expect(wr.Code).To(Equal(200))
			Expect(wr.Body.String()).To(Equal("5"))

			wr = httptest.NewRecorder()
			req = httptest.NewRequest("GET", "/add?a=1&b=100", nil)
			server.Routes().ServeHTTP(wr, req)

			Expect(wr.Code).To(Equal(200))
			Expect(wr.Body.String()).To(Equal("101"))
		})
	})

	Describe("/hellojson", func() {
		It("should return a JSON content type", func() {
			wr := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/hellojson", nil)
			server.Routes().ServeHTTP(wr, req)

			Expect(wr.Code).To(Equal(200))
			Expect(wr.Header().Get("Content-Type")).To(Equal("application/json"))
		})

		It("should return JSON data", func() {
			wr := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/hellojson", nil)
			server.Routes().ServeHTTP(wr, req)

			Expect(wr.Code).To(Equal(200))
			Expect(wr.Body.String()).To(Equal(`{"message": "Hello, world!"}`))
		})
	})
})
