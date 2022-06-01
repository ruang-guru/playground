package repository

import (
	"database/sql"
	"fmt"
)

type SalesRepository struct {
	db *sql.DB
}

func NewSalesRepository(db *sql.DB) *SalesRepository {
	return &SalesRepository{db: db}
}

func (u *SalesRepository) FetchSales(request GetSalesRequest) ([]Sales, error) {
	sqlStmt := `SELECT 
		s.id,
		s.product_id,
		s.quantity,
		s.date,
		p.product_name,
		p.price,
		p.category
	FROM 
		sales s
	INNER JOIN Products p ON s.product_id = p.id
	`

	if isValidRequest := request.IsValidRequest(); !isValidRequest {
		return nil, fmt.Errorf("bad request")
	}

	if request.StartPeriod != nil && request.EndPeriod != nil {
		sqlStmt = fmt.Sprintf("%s WHERE s.date BETWEEN ? AND ?", sqlStmt)
	}

	if request.ProductName != "" {
		sqlStmt = fmt.Sprintf("%s AND p.name = ?", sqlStmt)
	}

	rows, err := u.db.Query(sqlStmt, request.StartPeriod, request.EndPeriod, request.ProductName)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var sales []Sales

	for rows.Next() {
		var s Sales
		err := rows.Scan(&s.ID, &s.ProductID, &s.Quantity, &s.Date, &s.ProductName, &s.Price, &s.Category)
		if err != nil {
			return nil, err
		}
		s.Total = s.Price * s.Quantity
		sales = append(sales, s)
	}

	return sales, nil
}
