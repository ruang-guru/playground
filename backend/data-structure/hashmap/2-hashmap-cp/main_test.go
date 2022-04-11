package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	Describe("AnagramsChecker", func() {
		When("input two words with different lengths", func() {
			It("should return message `Bukan Anagram`", func() {
				anagram := AnagramsChecker("day", "friday")
				Expect(anagram).To(Equal("Bukan Anagram"))
			})
		})

		When("input two words arranged with the same letter composition", func() {
			It("should return message `Anagram`", func() {
				anagram := AnagramsChecker("fried", "fired")
				Expect(anagram).To(Equal("Anagram"))
			})
		})

		When("input two words arranged with different letter composition", func() {
			It("should return message `Bukan Anagram`", func() {
				anagram := AnagramsChecker("apple", "paddle")
				Expect(anagram).To(Equal("Bukan Anagram"))
			})
		})
	})
})
