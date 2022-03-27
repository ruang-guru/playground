package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Menu", func() {
	Describe("Get Menu", func() {
		It("returns slice of map string interface of menu", func() {
			menu := GetMenu()
			Expect(menu[0]["Nama"]).To(Equal("Ayam Goreng"))
			Expect(menu[0]["Jenis"]).To(Equal("Cepat saji"))
			Expect(menu[0]["Harga"]).To(Equal(20000))
			Expect(menu[1]["Nama"]).To(Equal("Cola"))
			Expect(menu[1]["Jenis"]).To(Equal("Minuman"))
			Expect(menu[1]["Harga"]).To(Equal(7000))
		})
	})
})
