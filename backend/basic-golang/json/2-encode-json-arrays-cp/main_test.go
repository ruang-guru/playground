package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	jsonencode "github.com/ruang-guru/playground/backend/basic-golang/json/2-encode-json-arrays-cp"
)

var _ = Describe("JSON Encode", func() {

	Describe("JSON Encode Array", func() {
		It("encoding struct array to string JSON array", func() {
			items := jsonencode.Items{[]jsonencode.Meja{
				{
					Jenis:     "Meja Lipat",
					Warna:     "Coklat",
					Jumlah:    40,
					Deskripsi: "meja untuk belajar",
				},
				{
					Jenis:     "Meja Hijau",
					Warna:     "Hijau",
					Jumlah:    10,
					Deskripsi: "meja untuk pengadilan",
				},
			}}
			meja := jsonencode.NewMeja(items)
			result := meja.EncodeJSON()
			Expect(result).To(Equal(`[{"jenis":"Meja Lipat","warna":"Coklat","jumlah":40,"deskripsi":"meja untuk belajar"},{"jenis":"Meja Hijau","warna":"Hijau","jumlah":10,"deskripsi":"meja untuk pengadilan"}]`))
		})
	})

})
