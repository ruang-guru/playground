package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	jsonencode "github.com/ruang-guru/playground/backend/basic-golang/json/1-encode-json-cp"
)

var _ = Describe("JSON Encode", func() {

	Describe("JSON Encode", func() {
		It("encoding struct to string JSON", func() {
			meja := jsonencode.NewMeja(jsonencode.Meja{Jenis: "Meja Belajar", Warna: "green", Jumlah: 2})
			result := meja.EncodeJSON()
			Expect(result).To(Equal(`{"jenis":"Meja Belajar","color":"green","jumlah":2}`))
		})
	})

})
