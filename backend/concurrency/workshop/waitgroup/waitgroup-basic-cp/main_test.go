package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("waitgroup", func() {
	It("can help to wait all goroutines to finish", func() {
		output := make(chan []bool)
		go testWG(output)
		done := <-output
		for i := 0; i < 5; i++ {
			Expect(done[i]).To(Equal(true))
		}
	})
})
