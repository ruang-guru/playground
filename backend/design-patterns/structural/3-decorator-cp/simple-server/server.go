package simpleserver

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Perlu kalian ketahui bahwa Interface atau signature yang akan kita gunakan disini adalah
// http.HandlerFunc https://pkg.go.dev/net/http#HandlerFunc dan http.Handler https://pkg.go.dev/net/http#Handler

// Decorator pattern di Go biasanya sering digunakan pada web server. Misalnya keperluan logging
type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func New() *Server {
	return &Server{}
}

type Server struct {
}

func (s *Server) GetPerson(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
}

type Logging struct {
}

// Karena agak ribet untuk melakukan testing pada stdout. Maka disini kita menggantinya dengan Header
func (l Logging) AddLogging(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) // TODO: replace this
}
