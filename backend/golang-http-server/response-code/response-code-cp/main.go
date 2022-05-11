package main

import (
	"fmt"
	"net/http"
)

var names = []string{
	"Tony",
	"Roger",
	"Bruce",
	"Stephen",
}

func IsNameExists(name string) bool {
	for _, n := range names {
		if n == name {
			return true
		}
	}

	return false
}

func GetNameHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: answer here
	}
}

func GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	// TODO: answer here
	return mux
}
