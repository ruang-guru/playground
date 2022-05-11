package main

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/ruang-guru/playground/backend/database/intro-to-sql/model"
)

var _ = Describe("SQL Create", func() {
	BeforeEach(func() {
		// Setup
		db, err := sql.Open("sqlite3", "./employee.db")
		if err != nil {
			panic(err)
		}

		_, err = db.Exec(`CREATE TABLE IF NOT EXISTS employees (
			id INTEGER PRIMARY KEY,
			nik VARCHAR(10) NOT NULL, 
			first_name VARCHAR(255) NOT NULL,
			last_name VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL
		);`)

		if err != nil {
			panic(err)
		}
	})

	AfterEach(func() {
		// Teardown
		db, err := sql.Open("sqlite3", "./employee.db")
		if err != nil {
			panic(err)
		}

		_, err = db.Exec(`DROP TABLE IF EXISTS employees;`)
		if err != nil {
			panic(err)
		}

		os.Remove("./employee.db")
	})

	Describe("Create", func() {
		It("should return last inserted ID", func() {
			db, err := sql.Open("sqlite3", "./employee.db")
			Expect(err).To(BeNil())

			employeeRepo := NewEmployeeRepository(db)

			employeeData := model.Employee{
				FirstName: "John",
				NIK:       "2022040199",
				LastName:  "Sunandar",
				Email:     "john.sunandar@gmail.com",
			}

			lastInsertedID, err := employeeRepo.CreateEmployee(&employeeData)
			Expect(err).To(BeNil())
			Expect(lastInsertedID).To(Equal(int64(1)))
		})
	})

})
