package main

// Syarat untuk menerapkan normalisasi bentuk kedua ini adalah data telah dibentuk dalam 1NF
// fungsi normalisasi 2NF antara lain :
// 1. Menghapus beberapa subset data yang ada pada tabel dan menempatkan mereka pada tabel terpisah.
// 2. Menciptakan hubungan antara tabel baru dan tabel lama dengan menciptakan foreign key.
// 3. Tidak ada atribut dalam tabel yang secara fungsional bergantung pada candidate key tabel tersebut.

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Vehicle struct {
	Id       int    `db:"id"`
	NoPolisi string `db:"no_polisi"`
	Warna    string `db:"warna"`
	Merek    string `db:"merek"`
	Tahun    string `db:"tahun"`
}

type Mekanik struct {
	Id          int    `db:"id"`
	MekanikID   string `db:"mekanik_id"`
	NamaMekanik string `db:"nama_mekanik"`
}

type Parts struct {
	Id        int    `db:"id"`
	KodeParts string `db:"kode_parts"`
	NamaParts string `db:"nama_parts"`
	Harga     int    `db:"harga"`
}

type Faktur struct {
	Id        int    `db:"id"`
	NoFaktur  int    `db:"no_faktur"`
	Tanggal   string `db:"tanggal"`
	NoPolisi  string `db:"no_polisi"`
	MekanikID string `db:"mekanik_id"`
	KodeParts string `db:"kode_parts"`
	Kuantitas int    `db:"kuantitas"`
	Harga     int    `db:"harga"`
	Discount  int    `db:"discount"`
	Potongan  int    `db:"potongan"`
}

// Migrate digunakan untuk melakukan migrasi database dengan data yang dibutuhkan
func Migrate() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./normalize.db")
	if err != nil {
		panic(err)
	}

	sqlStmt := `CREATE TABLE IF NOT EXISTS vehicle_2nf (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		no_polisi VARCHAR(20),
		warna VARCHAR(20),
		merek VARCHAR(20),
		tahun VARCHAR(20)
	);`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`
			INSERT INTO 
			vehicle_2nf (no_polisi, warna, merek, tahun)
			VALUES 
			    ("B3117lB", "Biru", "Supra X", "2020"),
				("B3117lB", "Biru", "Supra X", "2020");`)

	if err != nil {
		panic(err)
	}

	sqlStmt = `CREATE TABLE IF NOT EXISTS mekanik_2nf (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		mekanik_id VARCHAR(20),
		nama_mekanik VARCHAR(20)
	);`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
			INSERT INTO 
			mekanik_2nf (mekanik_id, nama_mekanik)
			VALUES 
			    ("DDE", "Djoko Dewanto");`)

	if err != nil {
		panic(err)
	}

	sqlStmt = `CREATE TABLE IF NOT EXISTS parts_2nf (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		kode_parts VARCHAR(20),
		nama_parts VARCHAR(20),
		harga INTEGER
	);`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
			INSERT INTO 
			parts_2nf (kode_parts, nama_parts, harga)
			VALUES 
			    ("20W501000CC", "Oli Top1 1000CC", 27000),
				("SERV001", "Engine Tune Up", 25000);`)

	if err != nil {
		panic(err)
	}

	sqlStmt = `CREATE TABLE IF NOT EXISTS faktur_2nf (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		no_faktur INTEGER,
		tanggal VARCHAR(20),
		no_polisi VARCHAR(20),
		mekanik_id VARCHAR(20),
		kode_parts VARCHAR(20),
		kuantitas INTEGER,
		harga INTEGER,
		discount INTEGER,
		potongan INTEGER
	);`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
			INSERT INTO 
			faktur_2nf (no_faktur, tanggal, no_polisi, mekanik_id, kode_parts, kuantitas, harga, discount, potongan)
			VALUES 
			    (05103214, "2020-01-01", "B3117lB", "DDE", "20W501000CC", 2, 27000, 1000, 2000),
				(05103214, "2020-01-01", "B3117lB", "DDE", "SERV001", 1, 25000, 2000, 2000);`)

	if err != nil {
		panic(err)
	}

	return db, nil
}

// Jalankan main untuk melakukan migrasi database
func main() {
	db, err := Migrate()
	if err != nil {
		panic(err)
	}

	// table 1
	rows, err := db.Query("SELECT * FROM vehicle_2nf")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var vehicle Vehicle
		err = rows.Scan(&vehicle.Id, &vehicle.NoPolisi, &vehicle.Warna, &vehicle.Merek, &vehicle.Tahun)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", vehicle)
	}

	// table 2
	rows, err = db.Query("SELECT * FROM mekanik_2nf")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var mekanik Mekanik
		err = rows.Scan(&mekanik.Id, &mekanik.MekanikID, &mekanik.NamaMekanik)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", mekanik)
	}

	// table 3
	rows, err = db.Query("SELECT * FROM parts_2nf")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var parts Parts
		err = rows.Scan(&parts.Id, &parts.KodeParts, &parts.NamaParts, &parts.Harga)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", parts)
	}

	// table 4
	rows, err = db.Query("SELECT * FROM faktur_2nf")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var faktur Faktur
		err = rows.Scan(&faktur.Id, &faktur.NoFaktur, &faktur.Tanggal, &faktur.NoPolisi, &faktur.MekanikID, &faktur.KodeParts, &faktur.Kuantitas, &faktur.Harga, &faktur.Discount, &faktur.Potongan)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", faktur)
	}

	defer rows.Close()
}
