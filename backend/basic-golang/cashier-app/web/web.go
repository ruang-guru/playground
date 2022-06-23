package web

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/ruang-guru/playground/backend/basic-golang/cashier-app/repository"
)

type WEB struct {
	usersRepo       repository.UserRepository
	productsRepo    repository.ProductRepository
	cartItemRepo    repository.CartItemRepository
	transactionRepo repository.TransactionRepository
	mux             *http.ServeMux
}

func NewWEB(usersRepo repository.UserRepository, productsRepo repository.ProductRepository, cartItemRepo repository.CartItemRepository, transactionRepo repository.TransactionRepository) WEB {
	mux := http.NewServeMux()
	web := WEB{
		usersRepo, productsRepo, cartItemRepo, transactionRepo, mux,
	}
	mux.HandleFunc("/login", web.loginPage)
	mux.HandleFunc("/signup", web.registerPage)
	mux.HandleFunc("/dashboard", web.DashboardPage)

	return web
}

func (web *WEB) BasePath() (*template.Template, error) {
	basePath, errPath := filepath.Abs("./web/views/*")
	if errPath != nil {
		return nil, errPath
	}
	var tmpl, err = template.ParseGlob(basePath)
	if err != nil {
		return nil, err
	}
	return tmpl, nil
}

func (web *WEB) Handler() *http.ServeMux {
	return web.mux
}

func (web *WEB) Start() {
	fmt.Println("starting web app at http://localhost:9000")
	http.ListenAndServe(":9000", web.Handler())
}
