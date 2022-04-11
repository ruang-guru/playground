package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Stack", func() {
	Describe("Peek on empty stack", func() {
		It("should return error empty stack", func() {
			stack := NewStack(10)
			_, err := stack.Peek()
			Expect(err).Should(HaveOccurred())
			Expect(err).To(Equal(ErrEmptyStack))
		})
	})

	Describe("Peek on stack with top equal to 1", func() {
		It("should return 1 and no error", func() {
			stack := NewStack(10)
			stack.Push(1)
			e, err := stack.Peek()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(e).To(Equal(1))
			Expect(stack.Top).To(Equal(0))
			Expect(stack.Data).To(HaveLen(1))
		})
	})
})
