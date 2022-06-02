package repository

import (
	"database/sql"
	"fmt"
	"time"
)

type TransactionRepository struct {
	db                 *sql.DB
	productsRepository ProductRepository
	cartItemRepository CartItemRepository
}

func NewTransactionRepository(db *sql.DB, productRepository ProductRepository, cartItemRepository CartItemRepository) TransactionRepository {
	return TransactionRepository{db, productRepository, cartItemRepository}
}

// Pay will take the cart items and add them to the sales table
// and then delete the cart items from the cart_items table
func (u *TransactionRepository) Pay(cartItems []CartItem, amount int) (int, error) {
	totalPrice, err := u.cartItemRepository.TotalPrice()
	if err != nil {
		return 0, err
	}

	if amount < totalPrice {
		return 0, fmt.Errorf("not enough money")
	}

	tx, err := u.db.Begin()
	if err != nil {
		return 0, err
	}

	for _, cartItem := range cartItems {
		product, err := u.productsRepository.FetchProductByID(cartItem.ProductID)
		if err != nil {
			tx.Rollback()
			return 0, err
		}

		if product.Quantity < cartItem.Quantity {
			tx.Rollback()
			return 0, err
		}

		// TODO: Implement SQL Query using the transaction connection to decrease current product's quantity
		// HINT: use `tx.Exec("... query ...")` instead of `u.db.Exec("... query ...")`
		// TODO: answer here

		// TODO: Implement SQL Query using the transaction connection to add cart item to sales table
		// HINT: use `tx.Exec("... query ...")` instead of `u.db.Exec("... query ...")`
		// TODO: answer here
	}

	// clear cart
	_, err = tx.Exec("DELETE FROM cart_items")
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	tx.Commit()

	moneyChanges := amount - totalPrice
	return moneyChanges, nil
}
