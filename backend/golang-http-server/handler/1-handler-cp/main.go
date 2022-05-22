package main

import (
	"fmt"
	"net/http"
	"time"
)

// Dari contoh yang diberikan, buatlah sebuah handler dengan menggunakan HandlerFunc yang menampilkan nama hari, bulan, dan tahun.
// Hint, gunakan time.Weekday, time.Day, time.Month, dan time.Year

func GetHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		now := time.Now()
		writer.Write([]byte(fmt.Sprintf("%v, %v %v %v", now.Weekday(), now.Day(), now.Month(), now.Year())))
	} // TODO: replace this
}
