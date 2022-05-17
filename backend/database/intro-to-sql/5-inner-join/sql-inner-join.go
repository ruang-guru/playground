package main

import (
	"database/sql"

	"github.com/ruang-guru/playground/backend/database/intro-to-sql/model"
)

type OrderRepo struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepo {
	return &OrderRepo{db}
}

func (r *OrderRepo) FetchOrders() ([]model.Order, error) {
	// statement dibawah ini adalah query yang akan dijalankan
	// menggunakan inner join
	// inner join adalah join yang menggabungkan 2 atau lebih tabel yang berhubungan
	// dengan menggunakan kondisi yang sama

	sqlStmt := `
	SELECT
		o.id
		, o.student_id
		, o.product_id
		, o.qty
		, o.order_date
		, s.name as student_name
		, p.name as product_name
		, p.price as product_price
	FROM orders o
	INNER JOIN students s ON o.student_id = s.id
	INNER JOIN products p ON o.product_id = p.id
	ORDER BY o.id`

	rows, err := r.db.Query(sqlStmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var orders []model.Order
	for rows.Next() {
		var o model.Order
		err := rows.Scan(
			&o.ID,
			&o.StudentID,
			&o.ProductID,
			&o.Quantity,
			&o.OrderDate,
			&o.StudentName,
			&o.ProductName,
			&o.ProductPrice,
		)
		if err != nil {
			return nil, err
		}
		o.TotalPrice = o.ProductPrice * o.Quantity
		orders = append(orders, o)
	}

	return orders, nil
}
