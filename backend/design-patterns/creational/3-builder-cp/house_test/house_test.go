package house_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ruang-guru/playground/backend/design-patterns/creational/3-builder-cp/house"
)

var _ = Describe("House", func() {
	Context("When zimbabwe builder is implemented", func() {
		It("Should be also use as a builder", func() {
			zw := house.NewBuilder("zimbabwe")
			zwKontraktor := house.NewKontraktor(zw)
			house := zwKontraktor.BuildHouse()
			Expect(house.NumOfWindows).To(Equal(2))
			Expect(house.HasGarage).To(BeTrue())
			Expect(house.HasSwimmingPool).To(BeTrue())
			Expect(house.NumOfDoors).To(Equal(1))
		})

		It("Should be also build a house without a swimmping pool", func() {
			zw := house.NewBuilder("zimbabwe")
			zwKontraktor := house.NewKontraktor(zw)
			house := zwKontraktor.BuildHouseWithoutSwimmingPool()
			Expect(house.NumOfWindows).To(Equal(1))
			Expect(house.HasGarage).To(BeTrue())
			Expect(house.HasSwimmingPool).To(BeFalse())
			Expect(house.NumOfDoors).To(Equal(1))
		})
	})
})
