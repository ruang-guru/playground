package main_test

import (
	"fmt"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Goroutine", func() {
	When("used with anonymous function", func() {
		It("can capture loop counter correctly", func() {
			result := 0
			for i := 0; i < 5; i++ {
				//jalankan goroutine yang melakukan print untuk tiap hasil perkalian i*i dalam loop
				//dan menambahkan hasil tersebut ke variable result
				// TODO: answer here

			}
			time.Sleep(10 * time.Millisecond)
			fmt.Println("main stop")
			Expect(result).To(Equal(30))
		})
	})
})

//pada output terdapat angka berikut:
// 4
// 0
// 16
// 1
// 9
// urutan tidak harus sesuai
