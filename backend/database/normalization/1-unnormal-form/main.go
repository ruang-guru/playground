package main

// perhatikan gambar unormal-example.png
// identifikasikan-lah semua attribut yang ada, lalu buatlah table seperti pada gambar  dengan nama table "unormal"

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Unormal struct {
	NoFaktur    int    `db:"no_faktur"`
	Tanggal     string `db:"tanggal"`
	NoPolisi    string `db:"no_polisi"`
	Warna       string `db:"warna"`
	Merek       string `db:"merek"`
	Tahun       string `db:"tahun"`
	MekanikId   string `db:"mekanik_id"`
	NamaMekanik string `db:"nama_mekanik"`
	KodeParts   string `db:"kode_parts"`
	NamaParts   string `db:"nama_parts"`
	Kuantitas   int    `db:"kuantitas"`
	Harga       int    `db:"harga"`
	Discount    int    `db:"discount"`
	Jumlah      int    `db:"jumlah"`
	Potongan    int    `db:"potongan"`
	Total       int    `db:"total"`
}

// Migrate digunakan untuk melakukan migrasi database dengan data yang dibutuhkan
func Migrate() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./normalize.db")
	if err != nil {
		panic(err)
	}

	sqlStmt := `CREATE TABLE IF NOT EXISTS unormal (
		no_faktur INTEGER,
		tanggal VARCHAR(12),
		no_polisi VARCHAR(10),
		warna VARCHAR(10),
		merek VARCHAR(10),
		tahun VARCHAR(10),
		mekanik_id VARCHAR(10),
		nama_mekanik VARCHAR(10),
		kode_parts VARCHAR(10),
		nama_parts VARCHAR(10),
		kuantitas INTEGER,
		harga INTEGER,
		discount INTEGER,
		jumlah INTEGER,
		potongan INTEGER,
		total INTEGER
	);`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`
			INSERT INTO 
			unormal (no_faktur, tanggal, no_polisi, warna, merek, tahun, mekanik_id, nama_mekanik, kode_parts, nama_parts, kuantitas, harga, discount, jumlah, potongan, total)
			VALUES 
			    (05103214, "2020-01-01", "B3117lB", "Biru", "Supra X", "2020", "DDE", "Djoko Dewanto", "20W501000CC", "Oli Top 1000cc", 2, 27000, 1000, 52000, 2000, 73000),
				(05103214, "2020-01-01", "B3117lB", "Biru", "Supra X", "2020", "DDE", "Djoko Dewanto", "SERV001", "Engine Tune Up", 1, 25000, 2000, 23000, 2000, 73000);`)

	if err != nil {
		fmt.Println("here")
	}

	return db, nil
}

func Rollback(db *sql.DB) {
	sqlStmt := `DROP TABLE unormal;`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		panic(err)
	}

}

// Jalankan main untuk melakukan migrasi database
func main() {
	db, err := Migrate()
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT * FROM unormal")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var unormal Unormal
		err = rows.Scan(&unormal.NoFaktur, &unormal.Tanggal, &unormal.NoPolisi, &unormal.Warna, &unormal.Merek, &unormal.Tahun, &unormal.MekanikId, &unormal.NamaMekanik, &unormal.KodeParts, &unormal.NamaParts, &unormal.Kuantitas, &unormal.Harga, &unormal.Discount, &unormal.Jumlah, &unormal.Potongan, &unormal.Total)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", unormal)
	}

	defer rows.Close()
}
