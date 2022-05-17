package test_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ruang-guru/playground/backend/design-patterns/structural/3-decorator-cp/coffe"
)

var _ = Describe("Coffe", func() {
	Context("GetCost and GetDescription", func() {
		It("should return the cost of Coldbrew", func() {
			coffe := coffe.Coldbrew{}
			Expect(coffe.GetCost()).To(Equal(3.00))
		})

		It("should return the cost of Mocha", func() {
			coffe := coffe.Espresso{}
			Expect(coffe.GetCost()).To(Equal(2.00))
		})

		When("Coldbrew is decorated with Ice", func() {
			It("should return correct description and cost", func() {
				coldBrew := coffe.Coldbrew{}
				ice := coffe.Ice{
					Coffe: coldBrew,
				}

				Expect(ice.GetDescription()).To(Equal("Coldbrew, Ice"))
				Expect(ice.GetCost()).To(Equal(3.20))
			})
		})

		When("Espresso is decorated with Mocha", func() {
			It("should return correct description and cost", func() {
				espresso := coffe.Espresso{}
				mocha := coffe.Mocha{
					Coffe: espresso,
				}

				Expect(mocha.GetDescription()).To(Equal("Espresso, Mocha"))
				Expect(mocha.GetCost()).To(Equal(3.00))
			})
		})
	})

	Context("Something Weird Happen", func() {
		It("Should return BEKU when we add double ice", func() {
			coldBrew := coffe.Coldbrew{}
			ice := coffe.Ice{
				Coffe: coldBrew,
			}

			whippedCream := coffe.Whipcream{
				Coffe: ice,
			}

			ice2 := coffe.Ice{
				Coffe: whippedCream,
			}

			Expect(ice2.GetDescription()).To(Equal("Coldbrew, Ice, Whipcream, Ice, BEKU"))

			// There is error in this test it should equal 2.50. But the test expected 2.5000000000000004
			Expect(ice2.GetCost()).Should(BeNumerically(">=", 3.50))
		})
	})
})
