package main

import (
	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/url-shortener/handlers"
	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/url-shortener/repository"
	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/url-shortener/router"
)

func main() {
	repo := repository.NewMapRepository()
	h := handlers.NewURLHandler(&repo)
	r := router.SetupRouter(h)

	r.Run()
}
