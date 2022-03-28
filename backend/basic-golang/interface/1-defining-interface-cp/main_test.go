package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Employee", func() {
	Describe("Get Manager Bonus", func() {
		It("returns 3 times base salary", func() {
			manager := Manager{
				BaseSalary: 20000000,
			}
			bonus := manager.GetBonus()
			Expect(bonus).To(Equal(3 * 20000000))
		})
	})
	Describe("Get Senior Engineer Bonus", func() {
		It("returns 2 times base salary", func() {
			seniorEngineer := SeniorEngineer{
				BaseSalary: 15000000,
			}
			bonus := seniorEngineer.GetBonus()
			Expect(bonus).To(Equal(2 * 15000000))
		})
	})
	Describe("Get Junior Engineer Bonus", func() {
		It("returns base salary", func() {
			juniorEngineer := JuniorEngineer{
				BaseSalary: 10000000,
			}
			bonus := juniorEngineer.GetBonus()
			Expect(bonus).To(Equal(10000000))
		})
	})
	Describe("Get Total Employee Bonus", func() {
		It("returns sum of employee bonus", func() {
			manager := Manager{
				BaseSalary: 20000000,
			}
			seniorEngineer := SeniorEngineer{
				BaseSalary: 15000000,
			}
			juniorEngineer := JuniorEngineer{
				BaseSalary: 10000000,
			}
			totalBonus := TotalEmployeeBonus([]Employee{manager, seniorEngineer, juniorEngineer})
			Expect(totalBonus).To(Equal(100000000))
		})
	})
})
