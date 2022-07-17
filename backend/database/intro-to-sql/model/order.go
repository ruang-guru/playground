package model

import "time"

type Order struct {
	ID           int       `db:"id"`
	StudentID    int       `db:"student_id"`
	StudentName  string    `db:"student_name"`
	ProductID    int       `db:"product_id"`
	ProductName  string    `db:"product_name"`
	ProductPrice int       `db:"product_price"`
	Quantity     int       `db:"qty"`
	OrderDate    time.Time `db:"order_date"`
	TotalPrice   int       `db:"total_price"`
}

type OrderStudent struct {
	ID          int       `db:"id"`
	StudentID   int       `db:"student_id"`
	StudentName *string   `db:"student_name"`
	Quantity    int       `db:"qty"`
	OrderDate   time.Time `db:"order_date"`
}

type OrderProduct struct {
	ID           int       `db:"id"`
	ProductID    int       `db:"product_id"`
	ProductName  *string   `db:"product_name"`
	ProductPrice *int      `db:"product_price"`
	Quantity     int       `db:"qty"`
	OrderDate    time.Time `db:"order_date"`
}
