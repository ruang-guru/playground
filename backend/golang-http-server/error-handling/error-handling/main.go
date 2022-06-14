package main

import (
	"fmt"
	"net/http"
)

// Error handling pada HTTP server tidak jauh beda dengan error handling golang pada umumnya.
// Hal yang menjadi tambahan pada error handling di HTTP server ini adalah kita perlu menentukan response code yang tepat untuk error terkait.
// Misalnya, ketika client request sebuah data tetapi data tersebut tidak ada, maka response code yang tepat adalah HTTP 404 Not Found.
// Untuk dapat return error, kita bisa menggunakan method http.Error. Pada method ini terdapat tiga parameter yaitu response writer, error message, dan error code.

var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprint(w, "Hello World")
}
