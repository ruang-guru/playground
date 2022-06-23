package main

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

//pada program ini kita akan menambahkan case ketika tidak ada movement atau attack input yang siap
var _ = Describe("channel", func() {
	When("using select", func() {
		It("can use default case", func() {
			john := character{
				name:            "john",
				defaultActivity: "mengawasi area",
				activity:        "",
			}

			movementInput := make(chan string, 4)
			attackInput := make(chan string, 3)

			go john.awake(movementInput, attackInput)
			movementInput <- "atas"
			attackInput <- "B"
			movementInput <- "kiri"
			attackInput <- "A"
			movementInput <- "kanan"
			attackInput <- "C"
			movementInput <- "bawah"
			time.Sleep(100 * time.Millisecond)
			Expect(john.activity).ToNot(Equal(""))
		})
	})
})
