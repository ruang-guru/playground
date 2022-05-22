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
	// fmt.Println(result)
	return result, nil
}

func (u *CartItemRepository) Save(cartItems []CartItem) error {
	u.db.Delete("cart_items")
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
	// datas, err := u.db.Load("cart_items")
	// if err != nil {
	// 	return nil, err
	// }
	// items := make([]CartItem, 1)
	// for _, val := range datas {
	// 	price, _ := strconv.Atoi(val[2])
	// 	qty, _ := strconv.Atoi(val[3])
	// 	items = append(items, CartItem{
	// 		Category:    val[0],
	// 		ProductName: val[1],
	// 		Price:       price,
	// 		Quantity:    qty,
	// 	})
	// }
	// return items, nil // TODO: replace this
	return u.LoadOrCreate()
}

func (u *CartItemRepository) Add(product Product) error {
	cartItems, err := u.LoadOrCreate()
	if err != nil {
		return err
	}

	isFound := false
	for i, val := range cartItems {

		if val.Category == product.Category && val.ProductName == product.ProductName && val.Price == product.Price {
			cartItems[i].Quantity = val.Quantity + 1
			isFound = true
			break
		}
	}

	if !isFound {
		cartItems = append(cartItems, CartItem{
			Category:    product.Category,
			ProductName: product.ProductName,
			Price:       product.Price,
			Quantity:    1,
		})
	}
	err = u.Save(cartItems)
	return err // TODO: replace this
}

func (u *CartItemRepository) ResetCartItems() error {

	return u.Save([]CartItem{}) // TODO: replace this
}

func (u *CartItemRepository) TotalPrice() (int, error) {
	cartItems, err := u.LoadOrCreate()
	if err != nil {
		return 0, err
	}

	total := 0

	for _, val := range cartItems {
		total += val.Price * val.Quantity
	}
	return total, nil // TODO: replace this
}
