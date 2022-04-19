package main

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	Describe("SearchMatch", func() {
		When("input two arrays are empty", func() {
			It("should return an error", func() {
				_, err := SearchMatch([]string{}, []string{})
				Expect(err).To(Equal(fmt.Errorf("no match")))
			})
		})
		When("input two arrays are not empty", func() {
			It("should return matched array", func() {
				var cars1 = []string{"Toyota", "Honda", "Nissan", "BMW", "Chevy", "Ford"}
				var cars2 = []string{"Ford", "BMW", "Audi", "Mercedes"}
				matchResult, _ := SearchMatch(cars1, cars2)
				Expect(matchResult).To(Equal([]string{"BMW", "Ford"}))
			})
		})
		When("input two arrays are not empty, but match is not found", func() {
			It("should return an error", func() {
				var cars1 = []string{"Toyota", "Honda", "Nissan", "Chevy", "Ford"}
				var cars2 = []string{"Audi", "Mercedes", "Tesla"}
				_, err := SearchMatch(cars1, cars2)
				Expect(err).To(Equal(fmt.Errorf("no match")))
			})
		})
	})
})
