package main

import (
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
	mainUI.Start()
}
