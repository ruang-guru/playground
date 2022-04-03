package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	Describe("ScanToArray(arr)", func() {
		It("Should return array/slice of string", func() {
			var arr []string
			fileName := "main.txt"
			err := ScanToArray(&arr, fileName)

			Expect(err).To(BeNil())
			Expect(arr).To(HaveLen(10))
			Expect(arr[0]).To(Equal("3"))
			Expect(arr[1]).To(Equal("1"))
			Expect(arr[2]).To(Equal("4"))
			Expect(arr[3]).To(Equal("1"))
			Expect(arr[4]).To(Equal("5"))
			Expect(arr[5]).To(Equal("9"))
			Expect(arr[6]).To(Equal("2"))
			Expect(arr[7]).To(Equal("6"))
			Expect(arr[8]).To(Equal("5"))
			Expect(arr[9]).To(Equal("3"))
		})
	})

	Describe("ScanToMap(map)", func() {
		It("Should return a map based on the txt file per line (key, value)", func() {
			mapData := make(map[string]string)
			fileName := "bookAuthor.txt"
			err := ScanToMap(mapData, fileName)
			Expect(err).To(BeNil())
			Expect(mapData).To(HaveLen(8))
			Expect(mapData["Epictetus"]).To(Equal("Enchiridion"))
			Expect(mapData["Henry Marampiring"]).To(Equal("Filosofi Teras"))
			Expect(mapData["Leonhard Euler"]).To(Equal("Elements of Algebra"))
			Expect(mapData["Isaac Newton"]).To(Equal("Principia"))
		})
	})
})
