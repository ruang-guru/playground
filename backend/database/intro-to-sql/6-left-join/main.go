package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/ruang-guru/playground/backend/database/intro-to-sql/helper"
)

func main() {
	db, err := sql.Open("sqlite3", "./example.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	orderRepo := NewOrderRepository(db)

	orderProducts, err := orderRepo.FetchOrderProducts()
	if err != nil {
		panic(err)
	}

	for _, orderProduct := range orderProducts {
		fmt.Printf("ID: %d, ProducID: %d, ProductName: %s, Quantity: %d, Price: %d, OrderDate: %s\n",
			orderProduct.ID, orderProduct.ProductID, helper.DereferenceString(orderProduct.ProductName), orderProduct.Quantity, helper.DereferenceInt(orderProduct.ProductPrice), orderProduct.OrderDate)
	}
}
