package main

import (
	"github.com/ruang-guru/playground/backend/basic-golang/cashier-app/api"
	"github.com/ruang-guru/playground/backend/basic-golang/cashier-app/db"
	"github.com/ruang-guru/playground/backend/basic-golang/cashier-app/repository"
	"github.com/ruang-guru/playground/backend/basic-golang/cashier-app/ui"
)

func main() {
	db := &db.CsvDB{}
	usersRepo := repository.NewUserRepository(db)
	productsRepo := repository.NewProductRepository(db)
	cartItemRepo := repository.NewCartItemRepository(db)
	transactionRepo := repository.NewTransactionRepository(cartItemRepo)

	mainUI := ui.NewUI(usersRepo, productsRepo, cartItemRepo, transactionRepo)
	//TODO: don't use go routine
	go mainUI.Start()

	mainAPI := api.NewAPI(usersRepo, productsRepo, cartItemRepo, transactionRepo)
	mainAPI.Start()
}
