package main

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/ruang-guru/playground/backend/database/intro-to-sql/model"
)

var _ = Describe("SQL Delete", func() {
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

	Describe("Delete employee data by NIK", func() {
		It("should success delete employee data", func() {
			db, err := sql.Open("sqlite3", "./employee.db")
			Expect(err).To(BeNil())

			employeeRepo := NewEmployeeRepository(db)
			err = employeeRepo.DeleteEmployeeByNIK("12345")
			Expect(err).To(BeNil())

			rows, err := db.Query("SELECT * FROM employees")
			Expect(err).To(BeNil())

			var employees []model.Employee
			for rows.Next() {
				var employee model.Employee
				err = rows.Scan(&employee.ID, &employee.NIK, &employee.FirstName, &employee.LastName, &employee.Email)
				Expect(err).To(BeNil())
				employees = append(employees, employee)
			}
			Expect(len(employees)).To(Equal(2))
		})
	})
})
