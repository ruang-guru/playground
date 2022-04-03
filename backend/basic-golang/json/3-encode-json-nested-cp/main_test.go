package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	jsonencode "github.com/ruang-guru/playground/backend/basic-golang/json/3-encode-json-nested-cp"
)

var _ = Describe("JSON Encode", func() {

	Describe("JSON Encode Array Nested", func() {
		It("encoding struct array to string JSON array", func() {
			items := jsonencode.Items{[]jsonencode.Meja{
				{
					Jenis:  "Meja Makan",
					Warna:  "Coklat",
					Jumlah: 20,
					Ukuran: jsonencode.Ukuran{
						Panjang: "50 cm",
						Tinggi:  "25 cm",
					},
				},
				{
					Jenis:  "Meja Lipat",
					Warna:  "Hitam",
					Jumlah: 1,
					Ukuran: jsonencode.Ukuran{
						Panjang: "70 cm",
						Tinggi:  "30 cm",
					},
				},
			}}
			meja := jsonencode.NewMeja(items)
			result := meja.EncodeJSON()
			Expect(result).To(Equal(`[{"jenis":"Meja Makan","warna":"Coklat","jumlah":20,"ukuran":{"panjang":"50 cm","tinggi":"25 cm"}},{"jenis":"Meja Lipat","warna":"Hitam","jumlah":1,"ukuran":{"panjang":"70 cm","tinggi":"30 cm"}}]`))
		})
	})

})
