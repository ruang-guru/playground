package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("SQL Open", func() {
	AfterEach(func() {
		db, err := sql.Open("sqlite3", "./studentData.db")
		if err != nil {
			panic(err)
		}
		defer db.Close()
	})

	Describe("Open ConnectSQLite()", func() {
		It("should return successfully open", func() {
			res, err := ConnectSQLite()
			Expect(err).To(BeNil())
			Expect(res).To(Equal("You are successfully opening the database studentData.db"))
		})
	})
})
