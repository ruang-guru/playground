package api

import (
	"encoding/json"
	"net/http"

	"github.com/ruang-guru/playground/backend/database/assigment/cashier-app/repository"
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

type PayRequest struct {
	Cash int `json:"cash"`
}

type PaySuccessResponse struct {
	Cash    int `json:"cash"`
	Total   int `json:"total_payment"`
	Changes int `json:"changes"`
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

	product, err := api.productsRepo.FetchProductByName(requestBody.ProductName)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		encoder.Encode(CartErrorResponse{Error: err.Error()})
		return
	}

	cart, err := api.cartItemRepo.FetchCartByProductID(product.ID)
	if err == nil && cart.ID != 0 {
		err = api.cartItemRepo.IncrementCartItemQuantity(cart)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			encoder.Encode(CartErrorResponse{Error: err.Error()})
			return
		}

		w.WriteHeader(http.StatusOK)
		encoder.Encode(AddToCartSuccessResponse{
			Name:     product.ProductName,
			Price:    product.Price,
			Category: product.Category,
		})
		return
	}

	err = api.cartItemRepo.InsertCartItem(repository.CartItem{
		ProductID: product.ID,
		Quantity:  1,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(CartErrorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	encoder.Encode(AddToCartSuccessResponse{
		Name:     product.ProductName,
		Price:    product.Price,
		Category: product.Category,
	})
}

func (api *API) clearCart(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	err := api.cartItemRepo.ResetCartItems()
	encoder := json.NewEncoder(w)
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(CartErrorResponse{Error: err.Error()})
		}
	}()

	w.WriteHeader(http.StatusOK)
}

func (api *API) cartList(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	cartItems, err := api.cartItemRepo.FetchCartItems()
	encoder := json.NewEncoder(w)
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(CartErrorResponse{Error: err.Error()})
		}
	}()

	encoder.Encode(CartListSuccessResponse{CartItems: cartItems})
}

func (api *API) pay(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)

	var requestBody PayRequest

	err := json.NewDecoder(req.Body).Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cartItems, err := api.cartItemRepo.FetchCartItems()
	encoder := json.NewEncoder(w)
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(CartErrorResponse{Error: err.Error()})
		}
	}()

	sumPrice, err := api.cartItemRepo.TotalPrice()
	if err != nil {
		return
	}

	moneyChanges, err := api.transactionRepo.Pay(cartItems, requestBody.Cash)
	if err != nil {
		return
	}

	encoder.Encode(PaySuccessResponse{
		Cash:    requestBody.Cash,
		Total:   sumPrice,
		Changes: moneyChanges,
	})
}
