package main_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	httpserver "github.com/ruang-guru/playground/backend/basic-golang/net-http/2-server-api-post-cp"
)

var _ = Describe("POST JSON Request", func() {

	Describe("POST JSON To Add table", func() {
		It("create request to add new table to the list", func() {
			tables := []httpserver.Table{
				{ID: "T012", Type: "Meja Rias", Color: "Coklat", Total: 3},
				{ID: "T011", Type: "Meja Kotak", Color: "Hitam", Total: 4},
			}
			tablesJSON, err := json.Marshal(tables)
			if err != nil {
				log.Fatalf("could not encode json: %v", err)
			}
			req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/tables", bytes.NewReader(tablesJSON))
			if err != nil {
				log.Fatalf("could not create request: %v", err)
			}

			rec := httptest.NewRecorder()
			httpserver.TablesHandler(rec, req)

			result := rec.Result()
			defer result.Body.Close()

			if result.StatusCode != http.StatusCreated {
				log.Fatalf("expected status created 201; got %v", result.StatusCode)
			}

			b, err := ioutil.ReadAll(result.Body)
			if err != nil {
				log.Fatalf("could not read response: %v", err)
			}

			Expect(string(b)).To(Equal(`{"status":"add tables succeed"}`))
		})
	})

})
