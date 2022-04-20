package main

import (
	"fmt"
	"time"

	helper "github.com/ruang-guru/playground/backend/concurrency/assignment/async-requests-cp/client/request"
)

var start time.Time

func init() {
	start = time.Now()
}

func main() {
	baseURL := "http://localhost:8080/"
	urls := []string{}
	for i := 0; i < 10; i++ {
		urls = append(urls, baseURL)
	}
	results := helper.AsyncHttpGets(urls)
	for _, result := range results {
		if result != nil && result.Body != nil {
			fmt.Printf("%s. Waktu %v\n", result.Status, time.Since(start))
		}
	}
}

//reference: https://matt.aimonetti.net/posts/2012-11-real-life-concurrency-in-go/
