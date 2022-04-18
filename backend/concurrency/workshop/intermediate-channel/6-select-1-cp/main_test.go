package main

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("channel", func() {
	When("using select", func() {
		It("can choose channel that ready first", func() {
			john := character{
				name:            "john",
				defaultActivity: "mengawasi area",
				activity:        "",
			}

			movementInput := make(chan string, 4)
			attackInput := make(chan string, 3)

			go john.awake(movementInput, attackInput)
			movementInput <- "atas"
			movementInput <- "bawah"
			attackInput <- "A"
			movementInput <- "kiri"
			attackInput <- "B"
			movementInput <- "kanan"
			attackInput <- "C"
			time.Sleep(100 * time.Millisecond)
			Expect(john.activity).ToNot(Equal(""))
		})
	})
})
