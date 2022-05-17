package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Student struct {
	ID   int    `db:"id"`
	NIM  string `db:"nim"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

type Product struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Price int    `db:"price"`
}

type Order struct {
	ID        int       `db:"id"`
	StudentID int       `db:"student_id"`
	ProductID int       `db:"product_id"`
	OrderDate time.Time `db:"order_date"`
	OrderQty  int       `db:"order_qty"`
}

// Migrate digunakan untuk melakukan migrasi database dengan data yang dibutuhkan
func Migrate() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./example.db")
	if err != nil {
		panic(err)
	}

	sqlStmt := `CREATE TABLE IF NOT EXISTS students (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		nim VARCHAR(10) NOT NULL,
		name TEXT,
		age INTEGER
	);`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
			INSERT INTO 
			    students (nim, name, age) 
			VALUES 
			    ("12345", "John", 20),
				("23456", "Jane", 21),
				("34567", "Jenny", 22);`)

	if err != nil {
		return nil, err
	}

	sqlStmt = `CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		price INTEGER
	);`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
			INSERT INTO 
			    products (name, price) 
			VALUES 
			    ("Laptop", 100000),
				("Mouse", 5000),
				("Keyboard", 10000);`)

	if err != nil {
		return nil, err
	}

	sqlStmt = `CREATE TABLE IF NOT EXISTS orders (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		student_id INTEGER,
		product_id INTEGER,
		order_date DATE,
		qty INTEGER,
		FOREIGN KEY (student_id) REFERENCES students(id),
		FOREIGN KEY (product_id) REFERENCES products(id)
	);`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
			INSERT INTO 
			    orders (student_id, product_id, order_date, qty)  
			VALUES 
			    (1, 1, '2018-10-20', 1),
				(1, 2, '2018-10-21', 2),
				(2, 3, '2018-10-22', 3),
				(3, 1, '2018-10-23', 4),
				(4, 1, '2018-10-24', 4),
				(3, 5, '2018-10-25', 4),
				(4, 1, '2018-10-26', 4);`)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func Rollback(db *sql.DB) {
	sqlStmt := `DROP TABLE students;`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		panic(err)
	}

	sqlStmt = `DROP TABLE products;`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		panic(err)
	}

	sqlStmt = `DROP TABLE orders;`
	_, err = db.Exec(sqlStmt)
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

	rows, err := db.Query("SELECT * FROM students")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var student Student
		err = rows.Scan(&student.ID, &student.NIM, &student.Name, &student.Age)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", student)
	}

	rows, err = db.Query("SELECT * FROM products")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var product Product
		err = rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", product)
	}

	rows, err = db.Query("SELECT * FROM orders")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var order Order
		err = rows.Scan(&order.ID, &order.StudentID, &order.ProductID, &order.OrderDate, &order.OrderQty)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", order)
	}

	defer rows.Close()

	// Use Rollback() to rollback the changes
	//Rollback(db)
}
