package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Stack", func() {
	Describe("Push element on empty stack", func() {
		It("should add the stack size and data by 1", func() {
			stack := NewStack(10)
			err := stack.Push(1)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(stack.Top).To(Equal(0))
			Expect(stack.Data).To(HaveLen(1))
			Expect(stack.Data[0]).To(Equal(1))
		})
	})
	Describe("Push element on stack with size 1", func() {
		It("should add the stack size and data by 1", func() {
			stack := NewStack(10)
			err := stack.Push(1)
			err = stack.Push(2)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(stack.Top).To(Equal(1))
			Expect(stack.Data).To(HaveLen(2))
			Expect(stack.Data[0]).To(Equal(1))
			Expect(stack.Data[1]).To(Equal(2))
		})
	})
	Describe("Push element on full size stack", func() {
		It("should return error stack overflow", func() {
			stack := NewStack(1)
			err := stack.Push(1)
			err = stack.Push(2)
			Expect(err).Should(HaveOccurred())
			Expect(err).To(Equal(ErrStackOverflow))
			Expect(stack.Top).To(Equal(0))
			Expect(stack.Data).To(HaveLen(1))
			Expect(stack.Data[0]).To(Equal(1))
		})
	})
})
