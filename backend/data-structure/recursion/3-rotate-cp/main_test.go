package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	Describe("Rotate", func() {
		When("input 0 index and list is not rotated", func() {
			It("should return rotated list first number to last number", func() {
				list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
				rotated := Rotate(0, list)
				Expect(rotated).To(Equal([]int{10, 2, 3, 4, 5, 6, 7, 8, 9, 1}))
			})
		})
	})
})
