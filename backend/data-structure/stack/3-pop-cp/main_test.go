package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Stack", func() {
	Describe("Pop on empty stack", func() {
		It("should return error stack underflow", func() {
			stack := NewStack(10)
			_, err := stack.Pop()
			Expect(err).Should(HaveOccurred())
			Expect(err).To(Equal(ErrStackUnderflow))
			Expect(stack.Top).To(Equal(-1))
			Expect(stack.Data).To(HaveLen(0))
		})
	})
	Describe("Pop on stack that contain 1 element", func() {
		It("should return the top element and no error", func() {
			stack := NewStack(10)
			stack.Push(1)
			e, err := stack.Pop()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(stack.Top).To(Equal(-1))
			Expect(stack.Data).To(HaveLen(0))
			Expect(e).To(Equal(1))
		})
	})
	Describe("Pop on stack that contain 2 element", func() {
		It("should return the top element and no error", func() {
			stack := NewStack(10)
			stack.Push(1)
			stack.Push(2)
			e, err := stack.Pop()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(stack.Top).To(Equal(0))
			Expect(stack.Data).To(HaveLen(1))
			Expect(e).To(Equal(2))
		})
	})
})
