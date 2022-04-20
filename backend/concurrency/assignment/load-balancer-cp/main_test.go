package main

import (
	"fmt"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Load balancer", func() {
	It("can spread work to worker", func() {
		in, out := make(chan *Work), make(chan *Work)
		for i := 0; i < nWorkers; i++ {
			go worker(in, out, i+1)
		}
		for i := 0; i < nRequesters; i++ {
			go createRequest(in, i+1)
		}

		//receiver
		go receiver(out)
		expected := map[string]bool{}
		for i := 0; i < 5; i++ {
			for j := 0; j < 3; j++ {
				data := fmt.Sprintf("worker %d accepted request from client %d", j+1, i+1)
				expected[data] = true
			}
		}
		time.Sleep(1000 * time.Millisecond)
		mu.Lock()
		Eventually(data).Should(Equal(expected))
		mu.Unlock()
		Expect(true).To(Equal(true))
	})
})
