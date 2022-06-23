package main

import (
	"fmt"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var start time.Time

func init() {
	start = time.Now()
}

var _ = Describe("Goroutine", func() {
	It("can stop loop concurrently", func() {
		loop := true
		go stopLoop(&loop)
		for loop {
			fmt.Println("loop")
			//check jika program sudah berjalan lebih dari 200ms
			if time.Since(start) > 200*time.Millisecond {
				break
			}
		}
		Expect(loop).To(Equal(false))
	})
})
