package repository

import "time"

type User struct {
	Username string
	Password string
	Role     string
	Loggedin bool
	Token    string
}

type Product struct {
	Category    string
	ProductName string
	Price       int
}

type CartItem struct {
	Category    string
	ProductName string
	Price       int
	Quantity    int
}

type Sales struct {
	Date        time.Time
	Category    string
	ProductName string
	Price       int
	Quantity    int
	Total       int
}

type GetSalesRequest struct {
	StartPeriod *time.Time
	EndPeriod   *time.Time
	ProductName string
}

func (r GetSalesRequest) IsEmptyRequest() bool {
	if r.StartPeriod == nil && r.EndPeriod == nil && r.ProductName == "" {
		return true
	}

	return false
}

func (r GetSalesRequest) IsValidRequest() bool {
	if r.StartPeriod == nil && r.EndPeriod != nil {
		return false
	}
	if r.StartPeriod != nil && r.EndPeriod == nil {
		return false
	}

	return true
}
