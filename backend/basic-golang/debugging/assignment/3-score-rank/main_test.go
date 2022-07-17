package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Score Rank", func() {
	Describe("A Rank", func() {
		It("Upper Limit", func() {
			res := ScoreRankCorrect(100)
			Expect(res).To(Equal("A"))
		})

		It("Mid", func() {
			res := ScoreRankCorrect(95)
			Expect(res).To(Equal("A"))
		})

		It("Lower Limit", func() {
			res := ScoreRankCorrect(90)
			Expect(res).To(Equal("A"))
		})
	})

	Describe("B Rank", func() {
		It("Upper Limit", func() {
			res := ScoreRankCorrect(89)
			Expect(res).To(Equal("B"))
		})

		It("Mid", func() {
			res := ScoreRankCorrect(85)
			Expect(res).To(Equal("B"))
		})

		It("Lower Limit", func() {
			res := ScoreRankCorrect(80)
			Expect(res).To(Equal("B"))
		})
	})

	Describe("C Rank", func() {
		It("Upper Limit", func() {
			res := ScoreRankCorrect(79)
			Expect(res).To(Equal("C"))
		})

		It("Mid", func() {
			res := ScoreRankCorrect(75)
			Expect(res).To(Equal("C"))
		})

		It("Lower Limit", func() {
			res := ScoreRankCorrect(70)
			Expect(res).To(Equal("C"))
		})
	})

	Describe("D Rank", func() {
		It("Upper Limit", func() {
			res := ScoreRankCorrect(69)
			Expect(res).To(Equal("D"))
		})

		It("Mid", func() {
			res := ScoreRankCorrect(65)
			Expect(res).To(Equal("D"))
		})

		It("Lower Limit", func() {
			res := ScoreRankCorrect(60)
			Expect(res).To(Equal("D"))
		})
	})

	Describe("E Rank", func() {
		It("Upper Limit", func() {
			res := ScoreRankCorrect(59)
			Expect(res).To(Equal("E"))
		})

		It("Mid", func() {
			res := ScoreRankCorrect(30)
			Expect(res).To(Equal("E"))
		})

		It("Lower Limit", func() {
			res := ScoreRankCorrect(0)
			Expect(res).To(Equal("E"))
		})
	})

	Describe("Outside Invalid", func() {
		It("Upper", func() {
			res := ScoreRankCorrect(120)
			Expect(res).To(Equal("INVALID"))
		})

		It("Lower", func() {
			res := ScoreRankCorrect(-4)
			Expect(res).To(Equal("INVALID"))
		})
	})
})
