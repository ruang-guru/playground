package main

import (
	"net/http"
)

func getRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Method is not allowed.", http.StatusMethodNotAllowed)
			return

		}

		next.ServeHTTP(w, r)
	})
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

func main() {
	handler := http.HandlerFunc(helloWorldHandler)
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: getRequestMiddleware(handler),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
