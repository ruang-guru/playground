package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./example.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	orderRepo := NewOrderRepository(db)

	orders, err := orderRepo.FetchOrders()
	if err != nil {
		panic(err)
	}

	// Print all orders
	for _, order := range orders {
		fmt.Printf("%+v\n", order)
	}
}
