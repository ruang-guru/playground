package main

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("semaphore", func() {
	It("can help to control throughput", func() {
		sayHello := func(name string) {
			//sayHello
		}
		worker := newWorker(sayHello, 5)
		worker.doWork("friend")
		time.Sleep(10 * time.Millisecond)
		Expect(worker.maxWork).To(Equal(5))
	})
})
