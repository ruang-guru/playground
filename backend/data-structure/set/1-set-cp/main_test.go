package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	Describe("Intersection", func() {
		When("input two lists with different lengths but have the same words", func() {
			It("should return list intersection with the same word", func() {
				var str1 = []string{"Java", "Python", "Javascript", "C ++", "C#"}
				var str2 = []string{"Swift", "Java", "Kotlin", "Python"}
				interception := Intersection(str1, str2)
				Expect(interception).To(HaveLen(2))
				Expect(interception).To(Equal([]string{"Java", "Python"}))
			})
		})
		When("input two lists with different lengths but not have the same words", func() {
			It("should return an empty list because it has no intersection with the same word", func() {
				var str1 = []string{"Java", "Python"}
				var str2 = []string{"Swift"}
				interception := Intersection(str1, str2)
				Expect(interception).To(HaveLen(0))
			})
		})
	})

	Describe("RemoveDuplicates", func() {
		When("input one lists that contains the same word", func() {
			It("will delete one of the same words, must return a list that has no duplicate words", func() {
				var str = []string{"Java", "Python", "Python", "Javascript", "C ++", "C#", "C#"}
				removeDuplicates := RemoveDuplicates(str)
				Expect(removeDuplicates).To(HaveLen(5))
				Expect(removeDuplicates).To(Equal([]string{"Java", "Python", "Javascript", "C ++", "C#"}))
			})
		})
	})
})
