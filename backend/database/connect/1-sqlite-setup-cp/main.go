package main

import (
	// replace with your own import
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Task :
// Buatlah kode koneksi untuk membuka database SQLite pada function connectSQLite().
// file database sudah tersedia pada direktori playground-internal/backend/database/1-sqlite-setup-cp/studentData.db

// bila ada error, silahkan return errornya
// bila berhasil silahkan return "You are successfully opening the database studentData.db"

func main() {
	res, err := ConnectSQLite()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(res)
}

func ConnectSQLite() (string, error) {
	sql.Open("", "") // replace with your own database path

	return "", nil // TODO: replace this
}
