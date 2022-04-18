package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Three Consecutive Character", func() {
	Describe("Exist", func() {
		It("Result in front", func() {
			res := ThreeConsecutiveCharacterPositionCorrect("www.ruangguru.com")
			Expect(res).To(Equal(1))
		})
	})

	Describe("Not Exist", func() {
		It("Partial Match middle", func() {
			res := ThreeConsecutiveCharacterPositionCorrect("helloworld")
			Expect(res).To(Equal(-1))
		})

		It("Partial Match in the End", func() {
			res := ThreeConsecutiveCharacterPositionCorrect("abcdefghijklmnopqrstuvwxyzz")
			Expect(res).To(Equal(-1))
		})
	})
})
