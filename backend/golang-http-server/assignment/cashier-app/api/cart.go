package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/cashier-app/repository"
)

type CartErrorResponse struct {
	Error string `json:"error"`
}

type AddToCartSuccessResponse struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Category string `json:"category"`
}

type AddToCardRequest struct {
	ProductName string `json:"product_name"`
}

type CartListSuccessResponse struct {
	CartItems []repository.CartItem `json:"cart_items"`
}

func (api *API) addToCart(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)

	var requestBody AddToCardRequest
	err := json.NewDecoder(req.Body).Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	encoder := json.NewEncoder(w)

	allProducts, err := api.productsRepo.SelectAll()
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(CartErrorResponse{Error: err.Error()})
		}
	}()

	log.Println(requestBody.ProductName)
	for _, product := range allProducts {
		if product.ProductName == requestBody.ProductName {
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
	api.AllowOrigin(w, req)
	err := api.cartItemRepo.ResetCartItems()
	encoder := json.NewEncoder(w)
	defer func() {
		if err != nil {
			// TODO: answer here
			encoder.Encode(CartErrorResponse{Error: err.Error()})
		}
	}()

	w.WriteHeader(http.StatusOK)
}

func (api *API) cartList(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	cartItems, err := api.cartItemRepo.SelectAll()
	encoder := json.NewEncoder(w)
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(CartErrorResponse{Error: err.Error()})
		}
	}()

	fmt.Println(cartItems)

	encoder.Encode(CartListSuccessResponse{CartItems: []repository.CartItem{}}) // TODO: replace this
}

func (api *API) pay(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)

	cartItems, err := api.cartItemRepo.SelectAll()
	encoder := json.NewEncoder(w)
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(CartErrorResponse{Error: err.Error()})
		}
	}()

	err = api.salesRepo.Add(cartItems)

	encoder.Encode(CartListSuccessResponse{CartItems: []repository.CartItem{}}) // TODO: replace this

}
