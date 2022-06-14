package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Run This Script for migration db
func main() {
	db, err := sql.Open("sqlite3", "backend/database/assigment/cashier-app/db/cashier-app.db")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
    id integer not null primary key AUTOINCREMENT,
    username varchar(255) not null,
    password varchar(255) not null,
    role varchar(255) not null,
    loggedin boolean not null
);

CREATE TABLE IF NOT EXISTS products (
    id integer not null primary key AUTOINCREMENT,
    product_name varchar(255) not null,
    category varchar(255) not null,
    price integer not null,
    quantity integer not null
);

CREATE TABLE IF NOT EXISTS cart_items (
    id integer not null primary key AUTOINCREMENT,
    product_id integer not null,
    quantity integer not null,
    FOREIGN KEY (product_id) REFERENCES products(id)
);

CREATE TABLE IF NOT EXISTS sales (
    id integer not null primary key AUTOINCREMENT,
    date DATE not null,
    product_id integer not null,
    quantity integer not null,
    FOREIGN KEY (product_id) REFERENCES products(id)
);

INSERT INTO users(username, password, role, loggedin) VALUES
    ('aditira', '1234', 'admin', false),
    ('dina', '4321', 'employee', false),
    ('dito', '2552', 'employee', false);

INSERT INTO products(product_name, category, price, quantity) VALUES
    ('Orange', 'Fruits', 5000, 100),
    ('Apple', 'Fruits', 2000, 100),
    ('Melon', 'Fruits', 4000, 100),
    ('Watermelon', 'Fruits', 10000, 100),
    ('Banana', 'Fruits', 4000, 100),
    ('Carrot', 'Vegetables', 2000, 100),
    ('Broccoli', 'Vegetables', 5200, 100),
    ('Cucumber', 'Vegetables', 3400, 100),
    ('Potato', 'Vegetables', 6500, 100),
    ('Tomato', 'Vegetables', 2200, 100),
    ('Coffee', 'Drink', 4300, 100),
    ('Milk', 'Drink', 4000, 100),
    ('Tea', 'Drink', 2700, 100);`)

	if err != nil {
		panic(err)
	}
}
