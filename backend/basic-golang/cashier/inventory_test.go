package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	main "github.com/ruang-guru/playground/backend/basic-golang/cashier"
)

var _ = Describe("Inventory", func() {

	Describe("Add", func() {
		Describe("adds item to inventory", func() {
			It("increase the inventory size", func() {
				sut := main.Inventory{}
				sut.Add(main.NewItem(1, "item", 1.0))
				sut.Add(main.NewItem(2, "item", 2.0))

				Expect(sut.Items()).To(HaveLen(2))
			})
		})
	})
})
