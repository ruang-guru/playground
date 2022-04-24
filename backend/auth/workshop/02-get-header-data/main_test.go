package main

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("GetHEADER", func() {
	It("should return Tokenmu adalah <cookie value>", func() {
		wr := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/welcome", nil)
		cookie := &http.Cookie{Name: "token", Value: "abcdef"}
		req.AddCookie(cookie)
		Routes().ServeHTTP(wr, req)

		Expect(wr.Code).To(Equal(200))
		Expect(wr.Body.String()).To(Equal("Tokenmu adalah abcdef!"))
	})
})
