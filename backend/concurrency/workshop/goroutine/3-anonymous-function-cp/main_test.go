package main

import (
	// TODO: answer here
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

//buat sebuah goroutine menggunakan anonymous function
//function print "selamat sore" dan ubah nilai called menjadi true
var _ = Describe("Goroutine", func() {
	It("can be used with anonymous function", func() {
		called := false
		// TODO: answer here
		time.Sleep(10 * time.Millisecond)
		Expect(called).To(Equal(true))
	})
})
