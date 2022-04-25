package main

import (
	"net/http"
	"net/http/httptest"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("JWT", func() {
	Describe("/signin", func() {
		It("should return jwt token when given correct credential", func() {
			wr := httptest.NewRecorder()

			bodyReader := strings.NewReader(`{"username": "user1", "password": "password1"}`)

			req := httptest.NewRequest("POST", "/signin", bodyReader)
			Routes().ServeHTTP(wr, req)

			Expect(wr.Code).To(Equal(200))

			cookies := wr.Result().Cookies()

			var isCookieTokenExist bool
			for _, c := range cookies {
				if c.Name == "token" {
					isCookieTokenExist = true
					break
				}
			}
			Expect(isCookieTokenExist).To(BeTrue())
		})

		It("should return non authorized when given wrong credential", func() {
			wr := httptest.NewRecorder()

			bodyReader := strings.NewReader(`{"username": "user2", "password": "password1"}`)

			req := httptest.NewRequest("POST", "/signin", bodyReader)
			Routes().ServeHTTP(wr, req)

			Expect(wr.Code).To(Equal(401))
		})
	})

	Describe("/welcome", func() {
		It("should return non authorized when request without token", func() {
			wr := httptest.NewRecorder()

			req := httptest.NewRequest("POST", "/welcome", nil)
			Routes().ServeHTTP(wr, req)

			Expect(wr.Code).To(Equal(401))
		})
	})

	Describe("/signin then /welcome", func() {
		It("should return authorized when request with correct token", func() {
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
			req2 := httptest.NewRequest("POST", "/welcome", nil)
			req2.AddCookie(cookie)

			Routes().ServeHTTP(wr2, req2)

			Expect(wr2.Code).To(Equal(200))
			Expect(wr2.Body.String()).To(Equal("Welcome user1!"))

		})
	})

})
