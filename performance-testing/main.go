package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ruangguru/playground/performance-testing/handlers"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/movies", handlers.GetMovies)
	router.HandleFunc("/movie", handlers.AddMovie)
	router.HandleFunc("/movie/{id}", handlers.GetMovie)
	router.HandleFunc("/", handlers.Home)

	log.Fatal(http.ListenAndServe(":8090", router))
}
