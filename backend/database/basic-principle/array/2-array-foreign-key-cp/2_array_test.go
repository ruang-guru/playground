package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	bdp "github.com/ruang-guru/playground/backend/database/basic-principle/array/2-array-foreign-key-cp"
)

var _ = Describe("Array Foreign Key", func() {
	Describe("foreign key student to school table equivalen to sql `SELECT * FROM student LEFT JOIN school ON student.SchoolID = school.ID;`", func() {
		It("should get school table by ShoolID from student table", func() {
			db1 := bdp.NewSchoolDB()
			db1.InsertSchool("SMA 1", "Jakarta")
			db1.InsertSchool("SMA 2", "Bandung")

			db2 := bdp.NewStudentDB()
			db2.InsertStudent("Gina", 21, 1)
			db2.InsertStudent("Rin", 22, 2)

			student := db2.WhereStudent(1)

			Expect(student.Name).To(Equal("Gina"))
			Expect(student.Age).To(Equal(21))
			Expect(student.SchoolID).To(Equal(1))

			school := db1.GetSchool(student.SchoolID)
			Expect(school.Name).To(Equal("SMA 1"))
			Expect(school.Address).To(Equal("Jakarta"))

			student = db2.WhereStudent(2)
			Expect(student.Name).To(Equal("Rin"))
			Expect(student.Age).To(Equal(22))
			Expect(student.SchoolID).To(Equal(2))

			school = db1.GetSchool(student.SchoolID)
			Expect(school.Name).To(Equal("SMA 2"))
			Expect(school.Address).To(Equal("Bandung"))
		})
	})
})
