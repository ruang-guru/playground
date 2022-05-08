package main

import (
	"fmt"
	"net/http"
	"time"
)

// Dari contoh yang telah diberikan, buatlah route untuk TimeHandler dan SayHelloHandler.
// Buatlah route "/time" pada TimeHandler dan "/hello" untuk SayHelloHandler dengan menggunakan http.HandleFunc

var TimeHandler = func(writer http.ResponseWriter, request *http.Request) {
	// TODO: answer here
	now := time.Now()
	writer.Write([]byte(fmt.Sprintf("%v, %v %v %v", now.Weekday(), now.Day(), now.Month(), now.Year())))
}

var SayHelloHandler = func(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
	name := r.URL.Query().Get("name")
	if name != "" {
		w.Write([]byte(fmt.Sprintf("Hello, %s!", name)))
	} else {
		w.Write([]byte("Hello there"))
	}
}

func main() {
	http.HandleFunc("/time", TimeHandler)
	http.HandleFunc("/hello", SayHelloHandler)
	http.ListenAndServe("localhost:8080", nil)
}
