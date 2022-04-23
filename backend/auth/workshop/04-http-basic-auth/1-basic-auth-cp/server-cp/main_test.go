package main

import (
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Server Basic Auth", func() {
	Describe("/login", func() {
		When("not using basic auth", func() {
			It("should return unauthorized", func() {
				wr := httptest.NewRecorder()
				req := httptest.NewRequest("GET", "/login", nil)
				Routes().ServeHTTP(wr, req)
				Expect(wr.Code).To(Equal(401))
				Expect(wr.Body.String()).To(Equal("Error parsing basic auth"))
			})
		})
		When("the username are incorrect", func() {
			It("should return unauthorized", func() {
				wr := httptest.NewRecorder()
				req := httptest.NewRequest("GET", "/login", nil)
				req.SetBasicAuth("unknow", "aditira123")
				Routes().ServeHTTP(wr, req)

				Expect(wr.Code).To(Equal(401))
				Expect(wr.Body.String()).To(Equal("{\"message\": \"Invalid username\"}"))
			})
		})
		When("the password are incorrect", func() {
			It("should return unauthorized", func() {
				wr := httptest.NewRecorder()
				req := httptest.NewRequest("GET", "/login", nil)
				req.SetBasicAuth("aditira", "unknow123")
				Routes().ServeHTTP(wr, req)

				Expect(wr.Code).To(Equal(401))
				Expect(wr.Body.String()).To(Equal("{\"message\": \"Invalid password\"}"))
			})
		})
		When("the username and password are correct", func() {
			It("should return a successful login response", func() {
				wr := httptest.NewRecorder()
				req := httptest.NewRequest("GET", "/login", nil)
				req.SetBasicAuth("aditira", "aditira123")
				Routes().ServeHTTP(wr, req)

				Expect(wr.Code).To(Equal(200))
				Expect(wr.Body.String()).To(Equal("{\"message\": \"welcome to CAMP 2022!\"}"))
			})
		})
	})
})
