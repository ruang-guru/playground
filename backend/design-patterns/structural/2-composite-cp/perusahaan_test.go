package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ruang-guru/playground/backend/design-patterns/structural/2-composite-cp/perusahaan"
)

var _ = Describe("Perusahaan", func() {
	When("VP have 2 subordinate (junior)", func() {
		It("Should return 40", func() {
			j1 := perusahaan.Junior{}
			j2 := perusahaan.Junior{}

			vp1 := perusahaan.VP{
				Subordinate: []perusahaan.Employee{j1, j2},
			}

			Expect(vp1.TotalDivisonSalary()).To(Equal(40))
		})
	})

	When("CTO Have 2 VP, VP-1 have 2 subordinate and VP2 have 1 subordinate", func() {
		It("Should return 80", func() {
			j1 := perusahaan.Junior{}
			j2 := perusahaan.Junior{}
			j3 := perusahaan.Junior{}

			vp1 := perusahaan.VP{
				Subordinate: []perusahaan.Employee{j1, j2},
			}

			vp2 := perusahaan.VP{
				Subordinate: []perusahaan.Employee{j3},
			}

			cto := perusahaan.CTO{
				Subordinate: []perusahaan.Employee{vp1, vp2},
			}

			Expect(cto.TotalDivisonSalary()).To(Equal(100))
		})
	})

	It("Can be defined that a CTO also dont have VP but direct subordinate (junior)", func() {
		j1 := perusahaan.Junior{}
		j2 := perusahaan.Junior{}

		cto := perusahaan.CTO{
			Subordinate: []perusahaan.Employee{j1, j2},
		}

		Expect(cto.TotalDivisonSalary()).To(Equal(50))
	})
})
