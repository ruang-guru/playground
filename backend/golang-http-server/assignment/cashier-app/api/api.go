package api

import (
	"fmt"
	"net/http"

	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/cashier-app/repository"
)

type API struct {
	usersRepo       repository.UserRepository
	productsRepo    repository.ProductRepository
	cartItemRepo    repository.CartItemRepository
	transactionRepo repository.TransactionRepository
	salesRepo       repository.SalesRepository
	mux             *http.ServeMux
}

func NewAPI(usersRepo repository.UserRepository, productsRepo repository.ProductRepository, cartItemRepo repository.CartItemRepository, transactionRepo repository.TransactionRepository, salesRepo repository.SalesRepository) API {
	mux := http.NewServeMux()
	api := API{
		usersRepo, productsRepo, cartItemRepo, transactionRepo, salesRepo, mux,
	}

	mux.Handle("/api/user/login", api.POST(http.HandlerFunc(api.login)))
	mux.Handle("/api/user/logout", api.POST(http.HandlerFunc(api.logout)))

	// API with AuthMiddleware:
	mux.Handle("/api/products", api.GET(api.AuthMiddleWare(http.HandlerFunc(api.productList))))
	mux.Handle("/api/cart/add", api.POST(api.AuthMiddleWare(http.HandlerFunc(api.addToCart))))
	mux.Handle("/api/cart/clear", api.GET(api.AuthMiddleWare(http.HandlerFunc(api.clearCart))))
	mux.Handle("/api/carts", api.GET(api.AuthMiddleWare(http.HandlerFunc(api.cartList))))
	mux.Handle("/api/dashboard", api.GET(api.AuthMiddleWare(http.HandlerFunc(api.dashboard))))

	// API with AuthMiddleware and AdminMiddleware
	mux.Handle("/api/admin/sales", api.GET(api.AuthMiddleWare(api.AdminMiddleware(http.HandlerFunc(api.getDashboard)))))

	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", api.Handler())
}
