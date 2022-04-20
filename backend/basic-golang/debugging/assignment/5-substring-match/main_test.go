package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Substring Match", func() {
	It("Match Found", func() {
		res := IsExistInSourceCorrect("hello", "ll")
		Expect(res).To(Equal(true))
	})

	It("Match Not Found", func() {
		res := IsExistInSourceCorrect("aaaa", "bb")
		Expect(res).To(Equal(false))
	})

	It("Match Not Found Bigger Search", func() {
		res := IsExistInSourceCorrect("aaaa", "bbbbbb")
		Expect(res).To(Equal(false))
	})

	It("Match Not Found Bigger Search Partial Match", func() {
		res := IsExistInSourceCorrect("bbaa", "aaaaaa")
		Expect(res).To(Equal(false))
	})

	It("Match Found Empty Search", func() {
		res := IsExistInSourceCorrect("a", "")
		Expect(res).To(Equal(true))
	})

	It("Exact Match", func() {
		res := IsExistInSourceCorrect("a", "a")
		Expect(res).To(Equal(true))
	})
})
