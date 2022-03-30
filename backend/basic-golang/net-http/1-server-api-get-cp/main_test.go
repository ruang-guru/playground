package main_test

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	httpserver "github.com/ruang-guru/playground/backend/basic-golang/net-http/1-server-api-get-cp"
)

var _ = Describe("GET JSON Response", func() {

	Describe("GET Table By total", func() {
		It("get json response from server by total table invalid request", func() {
			req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/table?total=", nil)
			if err != nil {
				log.Fatalf("could not create request: %v", err)
			}

			rec := httptest.NewRecorder()
			httpserver.TableHandler(rec, req)
			result := rec.Result()
			defer result.Body.Close()
			if result.StatusCode != http.StatusBadRequest {
				log.Fatalf("expected status bad request 400; got %v", result.StatusCode)
			}

			b, err := ioutil.ReadAll(result.Body)
			if err != nil {
				log.Fatalf("could not read response: %v", err)
			}

			Expect(string(b)).To(Equal("invalid total\n"))
		})
	})

	Describe("GET Table By total", func() {
		It("get json response from server by total table", func() {
			req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/table?total=3", nil)
			if err != nil {
				log.Fatalf("could not create request: %v", err)
			}

			rec := httptest.NewRecorder()
			httpserver.TableHandler(rec, req)
			result := rec.Result()
			defer result.Body.Close()
			if result.StatusCode != http.StatusOK {
				log.Fatalf("expected status OK; got %v", result.StatusCode)
			}

			b, err := ioutil.ReadAll(result.Body)
			if err != nil {
				log.Fatalf("could not read response: %v", err)
			}

			Expect(string(b)).To(Equal(`[{"id":"T001","type":"Meja Lipat","color":"Coklat","total":3},{"id":"T004","type":"Meja Hejau","color":"Hijau","total":3}]`))
		})
	})

	Describe("GET Table By total Not Found", func() {
		It("get json response from server by total and table not found", func() {
			req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/table?total=2", nil)
			if err != nil {
				log.Fatalf("could not create request: %v", err)
			}

			rec := httptest.NewRecorder()
			httpserver.TableHandler(rec, req)
			result := rec.Result()
			defer result.Body.Close()
			if result.StatusCode != http.StatusNotFound {
				log.Fatalf("expected status not found 404; got %v", result.StatusCode)
			}

			b, err := ioutil.ReadAll(result.Body)
			if err != nil {
				log.Fatalf("could not read response: %v", err)
			}

			Expect(string(b)).To(Equal(`{"status":"table not found"}` + "\n"))
		})
	})

})
