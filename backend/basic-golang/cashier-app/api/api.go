package api

import (
	"fmt"
	"net/http"

	"github.com/ruang-guru/playground/backend/basic-golang/cashier-app/repository"
)

type API struct {
	usersRepo       repository.UserRepository
	productsRepo    repository.ProductRepository
	cartItemRepo    repository.CartItemRepository
	transactionRepo repository.TransactionRepository
	mux             *http.ServeMux
}

func NewAPI(usersRepo repository.UserRepository, productsRepo repository.ProductRepository, cartItemRepo repository.CartItemRepository, transactionRepo repository.TransactionRepository) API {
	mux := http.NewServeMux()
	api := API{
		usersRepo, productsRepo, cartItemRepo, transactionRepo, mux,
	}
	mux.HandleFunc("/api/user/login", api.login)
	mux.HandleFunc("/api/user/logout", api.logout)
	mux.HandleFunc("/api/dashboard", api.dashboard)
	mux.HandleFunc("/api/products", api.productList)
	mux.HandleFunc("/api/cart/add", api.addToCart)
	// TODO: answer here

	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", api.Handler())
}
