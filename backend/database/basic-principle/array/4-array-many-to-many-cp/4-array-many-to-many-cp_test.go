package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	bdp "github.com/ruang-guru/playground/backend/database/basic-principle/array/4-array-many-to-many-cp"
)

var _ = Describe("Array Many to Many", func() {
	Describe("create UserPhone table with user and phone ID", func() {
		It("should get user and phone table by id from UserPhone", func() {
			db1 := bdp.NewPhoneDB()
			db1.InsertPhone(62, 1234567890)

			db2 := bdp.NewUserDB()
			db2.InsertUser("Gina", 20)

			db3 := bdp.NewUserPhoneDB()
			db3.InsertUserPhone(1, 1)

			for _, v := range db3 {
				Expect(v.UserID).To(Equal(1))
				Expect(db2.GetUser(v.UserID).Name).To(Equal("Gina"))
				Expect(v.PhoneID).To(Equal(1))
				Expect(db1.GetPhone(v.PhoneID).Number).To(Equal(1234567890))
			}
		})
	})
})
