package api

import (
	"encoding/json"
	"net/http"

	"github.com/ruang-guru/playground/backend/basic-golang/cashier-app/repository"
)

type CartErrorResponse struct {
	Error string `json:"error"`
}

type AddToCartSuccessResponse struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Category string `json:"category"`
}

type CartListSuccessResponse struct {
	CartItems []repository.CartItem `json:"cart_items"`
}

func (api *API) addToCart(w http.ResponseWriter, req *http.Request) {
	productName := req.URL.Query().Get("product_name")
	encoder := json.NewEncoder(w)

	allProducts, err := api.productsRepo.SelectAll()
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(CartErrorResponse{Error: err.Error()})
		}
	}()

	for _, product := range allProducts {
		if product.ProductName == productName {
			err = api.cartItemRepo.Add(product)
			if err != nil {
				return
			}
			encoder.Encode(AddToCartSuccessResponse{
				Name:     product.ProductName,
				Price:    product.Price,
				Category: product.Category,
			})
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func (api *API) clearCart(w http.ResponseWriter, req *http.Request) {
	// TODO: answer here
}

func (api *API) cartList(w http.ResponseWriter, req *http.Request) {
	encoder.Encode(CartListSuccessResponse{CartItems: []repository.CartItem{}}) // TODO: replace this
}
