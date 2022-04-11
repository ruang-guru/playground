package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	Describe("Sort", func() {
		When("the array is already sorted", func() {
			It("should return the same array", func() {
				var nums = []int{1, 2, 3, 4, 5}
				sort := Sort(nums)
				Expect(sort).To(Equal(nums))
			})
		})
		When("the array is not sorted", func() {
			It("should return the sorted array", func() {
				var nums = []int{4, 5, 2, 1, 3}
				sort := Sort(nums)
				Expect(sort).To(Equal([]int{1, 2, 3, 4, 5}))
			})
		})
	})

	Describe("RotateLeft", func() {
		When("the array is empty", func() {
			It("should return the same array", func() {
				var nums = []int{}
				rotateLeft := RotateLeft(0, nums)
				Expect(rotateLeft).To(Equal(nums))
			})
		})
		When("the array is not empty", func() {
			It("should return the rotated array", func() {
				var nums = []int{1, 2, 3, 4, 5}
				rotateLeft := RotateLeft(4, nums)
				Expect(rotateLeft).To(Equal([]int{5, 1, 2, 3, 4}))
			})
		})
	})
})
