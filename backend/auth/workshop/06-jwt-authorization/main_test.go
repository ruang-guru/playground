package main

import (
	"net/http"
	"net/http/httptest"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("JWTAuthorization", func() {
	Describe("/admin", func() {
		It("should return authorized when request with admin account", func() {
			wr1 := httptest.NewRecorder()

			bodyReader := strings.NewReader(`{"username": "user2", "password": "password2"}`)

			req1 := httptest.NewRequest("POST", "/signin", bodyReader)
			Routes().ServeHTTP(wr1, req1)

			Expect(wr1.Code).To(Equal(200))

			var cookie *http.Cookie
			for _, c := range wr1.Result().Cookies() {
				if c.Name == "token" {
					cookie = c
				}
			}

			wr2 := httptest.NewRecorder()
			req2 := httptest.NewRequest("POST", "/admin", nil)
			req2.AddCookie(cookie)

			Routes().ServeHTTP(wr2, req2)

			Expect(wr2.Code).To(Equal(200))
			Expect(wr2.Body.String()).To(Equal("Welcome Admin user2!"))

		})

		It("should return non authorized when request with student account", func() {
			wr1 := httptest.NewRecorder()

			bodyReader := strings.NewReader(`{"username": "user1", "password": "password1"}`)

			req1 := httptest.NewRequest("POST", "/signin", bodyReader)
			Routes().ServeHTTP(wr1, req1)

			Expect(wr1.Code).To(Equal(200))

			var cookie *http.Cookie
			for _, c := range wr1.Result().Cookies() {
				if c.Name == "token" {
					cookie = c
				}
			}

			wr2 := httptest.NewRecorder()
			req2 := httptest.NewRequest("POST", "/admin", nil)
			req2.AddCookie(cookie)

			Routes().ServeHTTP(wr2, req2)

			Expect(wr2.Code).To(Equal(401))

		})
	})

	Describe("/profile", func() {
		It("should return authorized when request with admin account", func() {
			wr1 := httptest.NewRecorder()

			bodyReader := strings.NewReader(`{"username": "user2", "password": "password2"}`)

			req1 := httptest.NewRequest("POST", "/signin", bodyReader)
			Routes().ServeHTTP(wr1, req1)

			Expect(wr1.Code).To(Equal(200))

			var cookie *http.Cookie
			for _, c := range wr1.Result().Cookies() {
				if c.Name == "token" {
					cookie = c
				}
			}

			wr2 := httptest.NewRecorder()
			req2 := httptest.NewRequest("POST", "/admin", nil)
			req2.AddCookie(cookie)

			Routes().ServeHTTP(wr2, req2)

			Expect(wr2.Code).To(Equal(200))
			Expect(wr2.Body.String()).To(Equal("Welcome Admin user2!"))

		})

		It("should return authorized when request with student account", func() {
			wr1 := httptest.NewRecorder()

			bodyReader := strings.NewReader(`{"username": "user1", "password": "password1"}`)

			req1 := httptest.NewRequest("POST", "/signin", bodyReader)
			Routes().ServeHTTP(wr1, req1)

			Expect(wr1.Code).To(Equal(200))

			var cookie *http.Cookie
			for _, c := range wr1.Result().Cookies() {
				if c.Name == "token" {
					cookie = c
				}
			}

			wr2 := httptest.NewRecorder()
			req2 := httptest.NewRequest("POST", "/admin", nil)
			req2.AddCookie(cookie)

			Routes().ServeHTTP(wr2, req2)

			Expect(wr2.Code).To(Equal(401))

		})
	})

})
