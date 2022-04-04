package api

import (
	"encoding/json"
	"net/http"
	"strconv"
)

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

func (api *API) dashboard(w http.ResponseWriter, req *http.Request) {
	encoder := json.NewEncoder(w)
	username, err := api.usersRepo.FindLoggedinUser()
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		encoder.Encode(DashboardErrorResponse{Error: err.Error()})
		return
	}

	cartItems, err := api.cartItemRepo.SelectAll()

	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(DashboardErrorResponse{Error: err.Error()})
			return
		}
	}()

	if err != nil {
		return
	}

	if err != nil {
		return
	}

	sumPrice, err := api.cartItemRepo.TotalPrice()

	if err != nil {
		return
	}

	cashParam := req.URL.Query().Get("cash")

	if cashParam == "" {
		cashParam = "0"
	}
	cash, err := strconv.Atoi(cashParam)

	if err != nil {
		return
	}

	moneyChanges, err := api.transactionRepo.Pay(cash)
	if err != nil {
		return
	}

	response := DashboardSuccessResponse{
		Username:     *username,
		TotalPrice:   sumPrice,
		MoneyChanges: moneyChanges,
	}
	for _, cartItem := range cartItems {
		response.CartItems = append(response.CartItems, CartItem{
			ProductName: cartItem.ProductName,
			Category:    cartItem.Category,
			Price:       cartItem.Price,
			Quantity:    cartItem.Quantity,
		})
	}
	encoder.Encode(response)
}
