package main

import (
	"fmt"
	"net/http"
)

var users = map[string]string{
	"aditira": "aditira123",
	"dito":    "123dito",
	"miya":    "12345",
}

func main() {
	fmt.Println("Starting Server at port :8080")
	http.ListenAndServe(":8080", Routes())
}

func isAuthorised(username, password string) bool {
	pass, ok := users[username]
	if !ok {
		return false
	}

	return password == pass
}

func Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		username, password, ok := r.BasicAuth()
		if !ok {
			w.Header().Add("WWW-Authenticate", `Basic realm="Give username and password"`)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"message": "No basic auth present"}`))
			return
		}

		if !isAuthorised(username, password) {
			w.Header().Add("WWW-Authenticate", `Basic realm="Give username and password"`)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"message": "Invalid username or password"}`))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "welcome to CAMP 2022!"}`))
	})

	return mux
}

//Encode auth aditira:aditira123 disini -> https://www.base64encode.org/

//$ curl -v -X GET http://localhost:8080/login -H "Authorization: Basic YWRpdGlyYTphZGl0aXJhMTIz"
