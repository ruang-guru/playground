package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// TODO: answer here
		if r.Method == "GET" {
			w.Write([]byte("Hello, world!"))
			return
		}
		http.Error(w, "", http.StatusMethodNotAllowed)
	})
	mux.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		// TODO: answer here
		if r.Method == "GET" {
			w.Write([]byte(r.URL.Query().Get("data")))
			return
		}
		http.Error(w, "", http.StatusMethodNotAllowed)
	})
	mux.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		// TODO: answer here
		if r.Method == "GET" {
			param := r.URL.Query()
			a, _ := strconv.Atoi(param.Get("a"))
			b, _ := strconv.Atoi(param.Get("b"))
			total := a + b
			w.Write([]byte(strconv.Itoa(total)))
			return
		}
		http.Error(w, "", http.StatusMethodNotAllowed)
	})
	mux.HandleFunc("/hellojson", func(w http.ResponseWriter, r *http.Request) {
		// TODO: answer here
		if r.Method == "GET" {

			resp := struct {
				Message string `json:"message"`
			}{
				Message: "Hello, world!",
			}

			bytesResponse, _ := json.Marshal(resp)

			w.Header().Add("Content-Type", "application/json")
			w.Write(bytesResponse)
			return
		}
		http.Error(w, "", http.StatusMethodNotAllowed)
	})

	return mux
}
func main() {
	http.ListenAndServe(":8080", Routes())
}
