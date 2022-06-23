package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/ruang-guru/playground/backend/database/intro-to-sql/model"
)

var _ = Describe("SQL Update", func() {
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

	})

	Describe("Update Employee Data", func() {
		It("should success update employee data", func() {
			db, err := sql.Open("sqlite3", "./employee.db")
			Expect(err).To(BeNil())

			employeeRepo := NewEmployeeRepository(db)

			updateEmployee := &model.Employee{
				NIK:       "12345",
				FirstName: "Robert",
				LastName:  "Doe",
				Email:     "robert.doe@gmail.com",
			}
			err = employeeRepo.UpdateEmployee(updateEmployee)
			Expect(err).To(BeNil())

			row, err := db.Query("SELECT * FROM employees WHERE nik = '12345'")
			Expect(err).To(BeNil())

			var employee model.Employee
			for row.Next() {
				err = row.Scan(&employee.ID, &employee.NIK, &employee.FirstName, &employee.LastName, &employee.Email)
				Expect(err).To(BeNil())
			}
			Expect(employee.NIK).To(Equal("12345"))
			Expect(employee.FirstName).To(Equal("Robert"))
			Expect(employee.LastName).To(Equal("Doe"))
			Expect(employee.Email).To(Equal("robert.doe@gmail.com"))
		})
	})

	Describe("Update Employee Email", func() {
		It("should success update employee email", func() {
			db, err := sql.Open("sqlite3", "./employee.db")
			Expect(err).To(BeNil())

			employeeRepo := NewEmployeeRepository(db)

			err = employeeRepo.UpdateEmployeeEmail("12345", "john_doe@gmail.com")
			Expect(err).To(BeNil())

			row, err := db.Query("SELECT * FROM employees WHERE nik = '12345'")
			Expect(err).To(BeNil())

			var employee model.Employee
			for row.Next() {
				err = row.Scan(&employee.ID, &employee.NIK, &employee.FirstName, &employee.LastName, &employee.Email)
				Expect(err).To(BeNil())
			}
			Expect(employee.NIK).To(Equal("12345"))
			Expect(employee.Email).To(Equal("john_doe@gmail.com"))
		})
	})
})
