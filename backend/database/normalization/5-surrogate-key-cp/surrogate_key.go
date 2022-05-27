package main

// pada contoh gambar surrogate_key, diketahui terdapat 2 table pendaftaran sekolah, dimana kedua table tersebut
// memilik kolom regirstration_no, name, dan percentage.
// gabunglah kedua table tersebut, dan buatlah surrogate key pada table yang baru.

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

//the struct are irrelevant for the code, but hint for column
type SchoolA struct {
	RegistrationNo int
	Name           string
	Percentage     float64
	Grade          string
	NationalRank   int
}

type SchoolB struct {
	RegistrationNo int
	Name           string
	Percentage     float64
}

type SurrogateTable struct {
	Id             int
	RegistrationNo int
	Name           string
	Percentage     float64
	Grade          string
	NationalRank   int
}

// Migrate digunakan untuk melakukan migrasi database dengan data yang dibutuhkan
// Tugas: Replace tanda ... dengan Query yang tepat pada fungsi Migrate:
func Migrate() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./surrogate.db")
	if err != nil {
		panic(err)
	}

	sqlStmt:= `CREATE TABLE ... ;` // TODO: replace this

	_, err = db.Exec(sqlStmt)
	if err != nil {
		fmt.Printf("%q: %s\n", err, sqlStmt)
	}

	_, err := db.Exec(`INSERT INTO school_a_cp .... ;`) // TODO: replace this

	if err != nil {
		fmt.Printf("%q: %s\n", err, sqlStmt)
	}

	sqlStmt = `CREATE TABLE....` // TODO: replace this

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`INSERT INTO school_b_cp .... ;`) // TODO: replace this

	if err != nil {
		panic(err)
	}

	sqlStmt = `CREATE TABLE ...` // TODO: replace this

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`INSERT INTO surrogate_table_cp .... ;`) // TODO: replace this

	if err != nil {
		panic(err)
	}

	return db, nil
}
