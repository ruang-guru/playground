package test_test

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	simpleserver "github.com/ruang-guru/playground/backend/design-patterns/structural/3-decorator-cp/simple-server"
)

var _ = Describe("Server", func() {
	It("should return a person", func() {
		req := httptest.NewRequest("GET", "/people", nil)
		w := httptest.NewRecorder()

		server := simpleserver.New()
		server.GetPerson(w, req)

		Expect(w.Code).To(Equal(http.StatusOK))
		Expect(w.Header().Get("Content-Type")).To(Equal("application/json"))
		Expect(w.Body.String()).To(MatchJSON(`{
			"name": "John Doe",
			"age": 30,
			"email": "john_doe@gmail.com"
		}`))
	})

	It("should return a person with logging header", func() {
		req := httptest.NewRequest("GET", "/people", nil)
		w := httptest.NewRecorder()

		server := simpleserver.New()

		log := simpleserver.Logging{}
		server2 := log.AddLogging(server.GetPerson)
		server2.ServeHTTP(w, req)

		Expect(w.Code).To(Equal(http.StatusOK))
		Expect(w.Header().Get("Content-Type")).To(Equal("application/json"))
		Expect(w.Header().Get("System-Log")).To(Equal("logged"))
		Expect(w.Body.String()).To(MatchJSON(`{
			"name": "John Doe",
			"age": 30,
			"email": "john_doe@gmail.com"
		}`))
	})

})
