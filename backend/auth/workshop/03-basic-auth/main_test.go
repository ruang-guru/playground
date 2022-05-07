package main

import (
	"net/http/httptest"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("BasicAuth", func() {
	It("should return authorized when given correct credential", func() {
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
