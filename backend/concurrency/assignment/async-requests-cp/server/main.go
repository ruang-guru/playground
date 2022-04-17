package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ruang-guru/playground/backend/concurrency/exercise/async-requests-cp/server/handler"
)

func main() {
	http.HandleFunc("/", handler.GetMessage)

	fmt.Println("starting server")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
