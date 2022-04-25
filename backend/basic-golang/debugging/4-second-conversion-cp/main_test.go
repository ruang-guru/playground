package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Second Conversion", func() {
	Describe("Limit 60", func() {
		It("Under One Minute", func() {
			res := ConvertSecondToTimeStringCorrect(30)
			Expect(res).To(Equal("00:00:30"))
		})

		It("Exact One Minute", func() {
			res := ConvertSecondToTimeStringCorrect(60)
			Expect(res).To(Equal("00:01:00"))
		})
	})

	Describe("Limit 3600", func() {
		It("Minute And Second", func() {
			res := ConvertSecondToTimeStringCorrect(70)
			Expect(res).To(Equal("00:01:10"))
		})

		It("Exact Hour", func() {
			res := ConvertSecondToTimeStringCorrect(3600)
			Expect(res).To(Equal("01:00:00"))
		})
	})

	Describe("More Than 3600", func() {
		It("Hour Minute And Second", func() {
			res := ConvertSecondToTimeStringCorrect(3670)
			Expect(res).To(Equal("01:01:10"))
		})

		It("Under One Day", func() {
			res := ConvertSecondToTimeStringCorrect(67812)
			Expect(res).To(Equal("18:50:12"))
		})

		It("Big Hour Minute And Second", func() {
			res := ConvertSecondToTimeStringCorrect(678120)
			Expect(res).To(Equal("188:22:00"))
		})
	})
})
