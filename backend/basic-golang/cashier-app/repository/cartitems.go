package repository

import (
	"strconv"

	"github.com/ruang-guru/playground/backend/basic-golang/cashier-app/db"
)

type CartItemRepository struct {
	db db.DB
}

func NewCartItemRepository(db db.DB) CartItemRepository {
	return CartItemRepository{db}
}

func (u *CartItemRepository) LoadOrCreate() ([]CartItem, error) {
	records, err := u.db.Load("cart_items")
	if err != nil {
		records = [][]string{
			{"category", "product_name", "price", "quantity"},
		}
		if err := u.db.Save("cart_items", records); err != nil {
			return nil, err
		}
	}

	result := make([]CartItem, 0)
	for i := 1; i < len(records); i++ {
		price, err := strconv.Atoi(records[i][2])
		if err != nil {
			return nil, err
		}

		qty, err := strconv.Atoi(records[i][3])
		if err != nil {
			return nil, err
		}

		cartItem := CartItem{
			Category:    records[i][0],
			ProductName: records[i][1],
			Price:       price,
			Quantity:    qty,
		}
		result = append(result, cartItem)
	}

	return result, nil
}

func (u *CartItemRepository) Save(cartItems []CartItem) error {
	records := [][]string{
		{"category", "product_name", "price", "quantity"},
	}
	for i := 0; i < len(cartItems); i++ {
		records = append(records, []string{
			cartItems[i].Category,
			cartItems[i].ProductName,
			strconv.Itoa(cartItems[i].Price),
			strconv.Itoa(cartItems[i].Quantity),
		})
	}
	return u.db.Save("cart_items", records)
}

func (u *CartItemRepository) SelectAll() ([]CartItem, error) {
	//beginanswer
	cartItems, err := u.LoadOrCreate()
	if err != nil {
		return nil, err
	}
	return cartItems, nil
	//endanswer return []CartItem{}, nil
}

func (u *CartItemRepository) Add(product Product) error {
	cartItems, err := u.LoadOrCreate()
	if err != nil {
		return err
	}

	//beginanswer

	//update
	for i := 0; i < len(cartItems); i++ {
		if cartItems[i].ProductName == product.ProductName {
			cartItems[i].Quantity++
			return u.Save(cartItems)
		}
	}

	//insert
	cartItems = append(cartItems, CartItem{
		Category:    product.Category,
		ProductName: product.ProductName,
		Price:       product.Price,
		Quantity:    1,
	})
	return u.Save(cartItems)
	//endanswer return nil
}

func (u *CartItemRepository) ResetCartItems() error {
	//beginanswer
	u.db.Delete("cart_items")
	records := [][]string{
		{"category", "product_name", "price", "quantity"},
	}
	return u.db.Save("cart_items", records)
	//endanswer return nil
}

func (u *CartItemRepository) TotalPrice() (int, error) {
	//beginanswer
	cartItems, err := u.LoadOrCreate()
	if err != nil {
		return 0, err
	}

	totalPrice := 0
	for i := 0; i < len(cartItems); i++ {
		totalPrice += cartItems[i].Price * cartItems[i].Quantity
	}

	return totalPrice, nil
	//endanswer return 0, nil
}
