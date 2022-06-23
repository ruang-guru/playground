package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	bdp "github.com/ruang-guru/playground/backend/database/basic-principle/array/1-array-cp"
)

var _ = Describe("Array", func() {
	Describe("update table by ID equivalent to sql `UPDATE table SET Name= 'John', Position= 'Senior Programmer', Salary= 150000000, ManagerID= 22  WHERE ID = 1;`", func() {
		It("should updated row table related by ID", func() {
			hm := bdp.NewEmployeeDB()
			hm.Insert("John", "Junior Programmer", 80000000, 22)
			hm.Update(1, "John", "Senior Programmer", 150000000, 22)

			val := hm.Where(1)
			Expect(val.Name).To(Equal("John"))
			Expect(val.Position).To(Equal("Senior Programmer"))
			Expect(val.Salary).To(Equal(150000000))
			Expect(val.ManagerID).To(Equal(22))
		})
	})

	Describe("delete table by ID equivalent to sql `DELETE FROM table WHERE ID = 1;`", func() {
		It("should deleted row table related by ID", func() {
			hm := bdp.NewEmployeeDB()
			hm.Insert("John", "Junior Programmer", 80000000, 22)
			hm.Insert("Gina", "Admin", 6000000, 18)
			hm.Delete(1)

			val := hm.Where(1)
			Expect(val).To(BeNil())

			val = hm.Where(2)
			Expect(val.Name).To(Equal("Gina"))
			Expect(val.Position).To(Equal("Admin"))
			Expect(val.Salary).To(Equal(6000000))
			Expect(val.ManagerID).To(Equal(18))
		})
	})
})
