package main

import (
	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/cashier-app/api"
	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/cashier-app/db"
	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/cashier-app/repository"
)

func main() {
	db := &db.CsvDB{}
	usersRepo := repository.NewUserRepository(db)
	productsRepo := repository.NewProductRepository(db)
	cartItemRepo := repository.NewCartItemRepository(db)
	salesRepo := repository.NewSalesRepository(db)
	transactionRepo := repository.NewTransactionRepository(cartItemRepo, salesRepo)

	mainAPI := api.NewAPI(usersRepo, productsRepo, cartItemRepo, transactionRepo, salesRepo)
	mainAPI.Start()

}
