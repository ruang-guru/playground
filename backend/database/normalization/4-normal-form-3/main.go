package main

// Normalisasi database dalam bentuk 3NF bertujuan untuk menghilangkan seluruh atribut atau field yang tidak berhubungan dengan primary key.
// Dengan demikian tidak ada ketergantungan transitif pada setiap kandidat key
// fungsi normalisasi 3NF antara lain :
// 1. Memenuhi semua persyaratan dari bentuk normal kedua.
// 2. Menghapus kolom yang tidak tergantung pada primary key.

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Member struct {
	IdMember   string `db:"id_member"`
	NamaMember string `db:"nama_member"`
	TipeMember string `db:"tipe_member"`
}

type TipeMember struct {
	TipeMember       string `db:"tipe_member"`
	KeteranganMember string `db:"keterangan_member"`
}

type Theater struct {
	IdTheater   string `db:"id_theater"`
	NamaTheater string `db:"nama_theater"`
	Kota        string `db:"kota"`
}

type Movie struct {
	IdMovie   string `db:"id_movie"`
	NamaMovie string `db:"nama_movie"`
}

type Faktur struct {
	IdFaktur    string `db:"id_faktur"`
	WaktuTayang string `db:"waktu_tayang"`
	HargaTiket  int    `db:"harga_tiket"`
	QtyTiket    int    `db:"qty_tiket"`
	IdMember    string `db:"id_member"`
	IdTheater   string `db:"id_theater"`
	IdMovie     string `db:"id_movie"`
}

// Migrate digunakan untuk melakukan migrasi database dengan data yang dibutuhkan
func Migrate() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./normalize.db")
	if err != nil {
		panic(err)
	}

	sqlStmt := `CREATE TABLE IF NOT EXISTS member_2nf (
		id_member VARCHAR(12) PRIMARY KEY,
		nama_member VARCHAR(12),
		tipe_member VARCHAR(10)
	);`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
	INSERT INTO 
	member_2nf (id_member, nama_member, tipe_member)
	VALUES
	("111", "Muri", "EPC"),
	("114", "Luga", "ELT");`)

	sqlStmt = `CREATE TABLE IF NOT EXISTS tipe_member_2nf (
		tipe_member VARCHAR(10) PRIMARY KEY,
		keterangan_member VARCHAR(10)
	);`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
	INSERT INTO 
	tipe_member_2nf (tipe_member, keterangan_member)
	VALUES
	("EPC", "Epic"),
	("ELT", "Elite");`)

	if err != nil {
		panic(err)
	}

	sqlStmt = `CREATE TABLE IF NOT EXISTS theater_2nf (
		id_theater VARCHAR(10) PRIMARY KEY,
		nama_theater VARCHAR(10),
		kota VARCHAR(10)
	);`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
			INSERT INTO 
			theater_2nf (id_theater, nama_theater, kota)
			VALUES 
			    ("T01", "Paris Van Java", "Bandung"),
			    ( "T02", "Grand Indonesia", "Jakarta");`)

	if err != nil {
		panic(err)
	}

	sqlStmt = `CREATE TABLE IF NOT EXISTS movie_2nf (
		id_movie VARCHAR(12) PRIMARY KEY,
		nama_movie VARCHAR(12)
	);`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
			INSERT INTO 
			movie_2nf (id_movie, nama_movie)
			VALUES 
			    ("M01", "Orang kaya baru"),
				("M03", "Twice Land"),
				("M04", "Escapee Room");`)

	if err != nil {
		panic(err)
	}

	sqlStmt = `CREATE TABLE IF NOT EXISTS faktur_2nf (
		id_faktur VARCHAR(12) PRIMARY KEY,
		waktu_tayang VARCHAR(12),
		harga_tiket INTEGER,
		qty_tiket INTEGER,
		id_member VARCHAR(12),
		id_theater VARCHAR(12),
		id_movie VARCHAR(12)
	);`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
			INSERT INTO 
			faktur_2nf (id_faktur, waktu_tayang, harga_tiket, qty_tiket, id_member, id_theater, id_movie)
			VALUES 
			    ("F001", "22/01/2022 15:00", "30000", "1","111","T01","M01"),
				("F002", "22/01/2022 18:00", "30000", "3","111","T01","M03"),
				("F003", "22/02/2022 15:00", "30000", "2","114","T02","M03"),
				("F004", "22/02/2022 18:00", "30000", "1","114","T02","M04");`)

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
	rows, err := db.Query("SELECT * FROM member_2nf")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var member Member
		err = rows.Scan(&member.IdMember, &member.NamaMember, &member.TipeMember)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", member)
	}

	// table 2
	rows, err = db.Query("SELECT * FROM tipe_member_2nf")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var tipeMember TipeMember
		err = rows.Scan(&tipeMember.TipeMember, &tipeMember.KeteranganMember)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", tipeMember)
	}

	// table 3
	rows, err = db.Query("SELECT * FROM theater_2nf")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var theater Theater
		err = rows.Scan(&theater.IdTheater, &theater.NamaTheater, &theater.Kota)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", theater)
	}

	// table 4
	rows, err = db.Query("SELECT * FROM movie_2nf")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var movie Movie
		err = rows.Scan(&movie.IdMovie, &movie.NamaMovie)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", movie)
	}

	// table 5
	rows, err = db.Query("SELECT * FROM faktur_2nf")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var faktur Faktur
		err = rows.Scan(&faktur.IdFaktur, &faktur.WaktuTayang, &faktur.HargaTiket, &faktur.QtyTiket, &faktur.IdMember, &faktur.IdTheater, &faktur.IdMovie)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", faktur)
	}

	defer rows.Close()
}
