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
	RegistrationNo string
	Name           string
	Percentage     float64
	Grade          string
	NationalRank   int
}

type SchoolB struct {
	RegistrationNo string
	Name           string
	Percentage     float64
}

type SurrogateTable struct {
	Id             int
	RegistrationNo string
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

	sqlStmt := `CREATE TABLE IF NOT EXISTS school_a_cp (
		registration_no Varchar(16),
		name TEXT,
		percentage REAL,
		grade INTEGER,
		national_rank INTEGER
	);`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		fmt.Printf("%q: %s\n", err, sqlStmt)
	}

	_, err = db.Exec(`
			INSERT INTO 
			school_a_cp (registration_no, name, percentage, grade, national_rank)
			VALUES 
			    ("sekolah1", "SMA Negeri 1", 0.5, "A", 457),
				("sekolah2", "SMA Negeri 2", 0.5, "A", 124),
				("sekolah3", "SMA Negeri 3", 0.5, "B", 789),
				("sekolah4", "SMA Negeri 4", 0.5, "B", 987),
				("sekolah5", "SMA Negeri 5", 0.5, "B", 1024);`)

	if err != nil {
		fmt.Printf("%q: %s\n", err, sqlStmt)
	}

	sqlStmt = `CREATE TABLE IF NOT EXISTS school_b_cp (
		registration_no Varchar(16),
		name TEXT,
		percentage REAL
	);`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
			INSERT INTO 
			school_b_cp (registration_no, name, percentage)
			VALUES 
			    ("1000", "SMA Negeri 6", 0.5),
				("2000", "SMA Negeri 7", 0.5),
				("3000", "SMA Negeri 8", 0.5),
				("4000", "SMA Negeri 9", 0.5),
				("5000", "SMA Negeri 10", 0.5);`)

	if err != nil {
		panic(err)
	}

	sqlStmt = `CREATE TABLE surrogate_table_cp ...` // TODO: replace this

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	//Masukkan data dua sekolah sebelumnya ke table ini
	_, err = db.Exec(`INSERT INTO surrogate_table_cp .... ;`) // TODO: replace this

	if err != nil {
		panic(err)
	}

	return db, nil
}
