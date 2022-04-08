package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	Describe("AgeDistribution", func() {
		When("input list person with name and age", func() {
			It("should return map with age and count", func() {
				people := []Person{{name: "Bob", age: 21}, {name: "Sam", age: 28}, {name: "Ann", age: 21}, {name: "Joe", age: 22}, {name: "Ben", age: 28}}
				Expect(AgeDistribution(people)).To(Equal(map[int]int{21: 2, 28: 2, 22: 1}))
			})
		})
	})

	Describe("FilterByAge", func() {
		When("input list person with name and age and filter by age", func() {
			It("should return list person name with age", func() {
				people := []Person{{name: "Bob", age: 21}, {name: "Sam", age: 28}, {name: "Ann", age: 21}, {name: "Joe", age: 22}, {name: "Ben", age: 28}}
				Expect(FilterByAge(people, 21)).To(Equal([]Person{{name: "Bob", age: 21}, {name: "Ann", age: 21}}))
			})
		})
	})
})
