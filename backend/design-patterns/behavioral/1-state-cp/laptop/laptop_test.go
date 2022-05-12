package laptop_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/ruang-guru/playground/backend/design-patterns/behavioral/1-state-cp/laptop"
)

var _ = Describe("Laptop", func() {
	Describe("Laptop sleeping behavior", func() {
		When("Laptop is on and there is battery left", func() {
			It("Should be able to sleep when laptop is off", func() {
				// laptop is off
				pc := laptop.New()
				// turn laptop on
				pc.Press()
				// pc.State
				pc.Sleep()
				Expect(pc.CurrentState).Should(Equal("Sleeping"))
			})
		})

		When("Laptop is off", func() {
			It("Should not be able to sleep", func() {
				pc := laptop.New()
				pc.Sleep()
				Expect(pc.CurrentState).Should(Equal("Off"))
			})
		})
	})
})
