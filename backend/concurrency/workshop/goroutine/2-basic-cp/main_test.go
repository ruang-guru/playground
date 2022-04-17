package main

import (

	// TODO: answer here

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Goroutine", func() {
	When("blocking happen", func() {
		It("switch to other goroutine", func() {
			multiplyCalled := false
			start(&multiplyCalled)
			Expect(multiplyCalled).To(Equal(true))
		})
	})
})
