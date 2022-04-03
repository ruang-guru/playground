package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	main "github.com/ruang-guru/playground/backend/golang-io/assignment/csv-file/read-csv-file-cp"
)

var _ = Describe("ReadUsers", func() {
	It("read the CSV file and parse it into User struct", func() {
		users, err := main.ReadUsers("./users.csv")
		Expect(err).To(Not(HaveOccurred()))

		Expect(users).To(HaveLen(3))
		Expect(users[0].Name).To(Equal("levi"))
		Expect(users[0].Age).To(Equal(32))
		Expect(users[0].Country).To(Equal("Indonesia"))
	})
})
