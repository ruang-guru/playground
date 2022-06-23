package api

type DashboardErrorResponse struct {
	Error string `json:"error"`
}

type CartItem struct {
	ProductName string `json:"product_name"`
	Category    string `json:"category"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
}

type DashboardSuccessResponse struct {
	Username     string     `json:"username"`
	CartItems    []CartItem `json:"cart_items"`
	TotalPrice   int        `json:"total_price"`
	MoneyChanges int        `json:"money_changes"`
}
