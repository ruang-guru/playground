package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	bdp "github.com/ruang-guru/playground/backend/database/basic-principle/array/3-array-one-to-many-cp"
)

var _ = Describe("Array One to Many", func() {
	Describe("join phone to student table equivalen to sql `SELECT * FROM phone LEFT JOIN student ON phone.StudentID = student.ID;`", func() {
		It("should get phone table by StudentID from student table", func() {
			db1 := bdp.NewPhoneDB()
			db1.InsertPhone(62, 1234567890, 1)
			db1.InsertPhone(62, 2345678901, 1)
			db1.InsertPhone(62, 3243434343, 1)

			db2 := bdp.NewUserDB()
			db2.InsertUser("John", 20)
			db2.InsertUser("Gina", 20)

			phone := db1.WherePhone(3)
			Expect(phone.ID).To(Equal(3))
			Expect(phone.CountryCode).To(Equal(62))
			Expect(phone.Number).To(Equal(3243434343))

			user := db2.GetUser(phone.UserID)
			Expect(user.ID).To(Equal(1))
			Expect(user.Name).To(Equal("John"))
			Expect(user.Age).To(Equal(20))
		})
	})
})
