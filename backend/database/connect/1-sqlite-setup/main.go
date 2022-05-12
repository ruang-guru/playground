package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		fmt.Println("error opening database:", err)
	}

	fmt.Println("You are successfully opening the database.")
	defer db.Close()
}