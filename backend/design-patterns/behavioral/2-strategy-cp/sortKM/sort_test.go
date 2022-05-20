package sortKM_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/ruang-guru/playground/backend/design-patterns/behavioral/2-strategy-cp/sortKM"
)

var _ = Describe("Sort strategy pattern", func() {
	When("strategy is AscendingSort", func() {
		It("will sort by ascending order", func() {
			array := []int{4, 5, 2, 1, 2, 6, 9}
			Sort := &SortKM{}
			Sort.SetStrategy(&AscendingSort{})
			Sort.Sort(array)
			fmt.Println(array)

			res := []int{1, 2, 2, 4, 5, 6, 9}
			Expect(array).To(Equal(res))
		})
	})
	When("strategy is DescendingSort", func() {
		It("will sort by descending order", func() {
			array := []int{4, 5, 2, 1, 2, 6, 9}
			Sort := &SortKM{}
			Sort.SetStrategy(&DescendingSort{})
			Sort.Sort(array)
			fmt.Println(array)

			res := []int{9, 6, 5, 4, 2, 2, 1}
			Expect(array).To(Equal(res))
		})
	})
})
