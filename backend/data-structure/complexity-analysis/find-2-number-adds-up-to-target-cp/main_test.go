package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	Describe("Two Target Sums", func() {
		When("found the position of two numbers that match the target if added up", func() {
			It("should return the two numbers position", func() {
				result := TwoTargetSums([]int{2, 5, 1, 3}, 6)
				Expect(result).To(Equal([]int{1, 2}))
			})
		})
		When("not found the position of two numbers that match the target if added up", func() {
			It("should return the two numbers position", func() {
				result := TwoTargetSums([]int{2, 5, 1, 3}, 100)
				Expect(result).To(Equal([]int{0, 0}))
			})
		})
	})
})
