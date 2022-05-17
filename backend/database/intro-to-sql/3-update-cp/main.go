package main

import (
	"database/sql"

	"github.com/ruang-guru/playground/backend/database/intro-to-sql/model"
)

type EmployeeRepository struct {
	db *sql.DB
}

func NewEmployeeRepository(db *sql.DB) *EmployeeRepository {
	return &EmployeeRepository{db}
}

func (r *EmployeeRepository) UpdateEmployee(updateEmployee *model.Employee) error {
	var sqlStmt string

	// Task : Create SQL statement
	// 1. Update the employee with the given id
	// 2. Set the updated fields (first_name, last_name, email)
	// TODO: answer here

	_, err := r.db.Exec(sqlStmt, updateEmployee.FirstName, updateEmployee.LastName, updateEmployee.Email, updateEmployee.NIK)
	if err != nil {
		return err
	}

	return nil
}

func (r *EmployeeRepository) UpdateEmployeeEmail(nik string, email string) error {
	var sqlStmt string

	// Task : Create SQL statement
	// 1. Update the employee with the given nik
	// 2. Set the updated fields (email)
	
	// TODO: answer here

	_, err := r.db.Exec(sqlStmt, email, nik)
	if err != nil {
		return err
	}

	return nil
}
