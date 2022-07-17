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

func (r *OrderRepo) FetchOrderProducts() ([]model.OrderProduct, error) {
	// statement berikut akan mengambil data dari tabel orders
	// menggunakan left join dengan tabel products
	// left join akan menghasilkan data yang kosong jika tidak ada data di tabel products

	sqlStmt := `
	SELECT
		o.id,
		o.product_id,
		p.name AS product_name,
		p.price AS product_price,
		o.qty,
		o.order_date
	FROM
		orders o
	LEFT JOIN
		products p
	ON
		o.product_id = p.id`

	rows, err := r.db.Query(sqlStmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var orders []model.OrderProduct
	for rows.Next() {
		var order model.OrderProduct
		err := rows.Scan(
			&order.ID,
			&order.ProductID,
			&order.ProductName,
			&order.ProductPrice,
			&order.Quantity,
			&order.OrderDate,
		)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}
