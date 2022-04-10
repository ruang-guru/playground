package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Stack", func() {
	Describe("Init stack with size 10", func() {
		It("should have stack with size equal to 10", func() {
			stack := NewStack(10)
			Expect(stack.Size).To(Equal(10))
			Expect(stack.Top).To(Equal(-1))
			Expect(stack.Data).To(HaveLen(0))
		})
	})
})
