package main

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/ruang-guru/playground/backend/database/intro-to-sql/model"
)

var _ = Describe("SQL Read", func() {
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

		_, err = db.Exec(`INSERT INTO employees (nik, first_name, last_name, email) VALUES
			("12345", "John", "Doe", "john.doe@gmail.com"),
			("54321", "Jane", "Doe", "jane.doe@gmail.com"),
			("98765", "John", "Smith", "john.smith@gmail.com");`)

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

	employeesData := []model.Employee{
		{
			ID:        1,
			NIK:       "12345",
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john.doe@gmail.com",
		},
		{
			ID:        2,
			NIK:       "54321",
			FirstName: "Jane",
			LastName:  "Doe",
			Email:     "jane.doe@gmail.com",
		},
		{
			ID:        3,
			NIK:       "98765",
			FirstName: "John",
			LastName:  "Smith",
			Email:     "john.smith@gmail.com",
		},
	}

	Describe("Fetch Employee by NIK", func() {
		It("should return employee data", func() {
			db, err := sql.Open("sqlite3", "./employee.db")
			Expect(err).To(BeNil())

			employeeRepo := NewEmployeeRepository(db)
			employee, err := employeeRepo.FetchEmployeeByNIK("12345")
			Expect(err).To(BeNil())
			Expect(employee).To(Equal(&employeesData[0]))
		})
	})

	Describe("Fetch Employees", func() {
		It("should return all employees data", func() {
			db, err := sql.Open("sqlite3", "./employee.db")
			Expect(err).To(BeNil())

			employeeRepo := NewEmployeeRepository(db)
			employees, err := employeeRepo.FetchEmployees()
			Expect(err).To(BeNil())
			Expect(employees).To(Equal(employeesData))
		})
	})
})
