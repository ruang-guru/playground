package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	Describe("Tower of Hanoi", func() {
		When("input number of disk and tower A, B and C", func() {
			It("should print solution", func() {
				n := 4
				TowerOfHanoi(n, "A", "B", "C")
				Expect(Solution[0]).To(Equal("Move disk 1 from rod A to rod C"))
				Expect(Solution[1]).To(Equal("Move disk 2 from rod A to rod C"))
				Expect(Solution[2]).To(Equal("Move disk 1 from rod B to rod A"))
				Expect(Solution[3]).To(Equal("Move disk 3 from rod A to rod C"))
				Expect(Solution[4]).To(Equal("Move disk 1 from rod B to rod A"))
				Expect(Solution[5]).To(Equal("Move disk 2 from rod B to rod A"))
				Expect(Solution[6]).To(Equal("Move disk 1 from rod C to rod B"))
				Expect(Solution[7]).To(Equal("Move disk 4 from rod A to rod C"))
				Expect(Solution[8]).To(Equal("Move disk 1 from rod B to rod A"))
				Expect(Solution[9]).To(Equal("Move disk 2 from rod B to rod A"))
				Expect(Solution[10]).To(Equal("Move disk 1 from rod C to rod B"))
				Expect(Solution[11]).To(Equal("Move disk 3 from rod B to rod A"))
				Expect(Solution[12]).To(Equal("Move disk 1 from rod C to rod B"))
				Expect(Solution[13]).To(Equal("Move disk 2 from rod C to rod B"))
				Expect(Solution[14]).To(Equal("Move disk 1 from rod A to rod C"))

			})
		})
	})
})
