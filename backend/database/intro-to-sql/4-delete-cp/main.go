package main

import "database/sql"

type EmployeeRepository struct {
	db *sql.DB
}

func NewEmployeeRepository(db *sql.DB) *EmployeeRepository {
	return &EmployeeRepository{db}
}

func (r *EmployeeRepository) DeleteEmployeeByNIK(nik string) error {
	var sqlStmt string

	// Task: delete employee by nik
	// 1. create sql statement
	// 2. use nik to delete employee

	// TODO: answer here

	_, err := r.db.Exec(sqlStmt, nik)
	if err != nil {
		return err
	}

	return nil
}
