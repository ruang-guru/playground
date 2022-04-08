package terminal

import (
	"fmt"
)

// ProductListPage is a function to show product list
func (terminal *Terminal) ProductPage() error {
	fmt.Printf("\x1bc") // clear screen

	products, err := terminal.productsRepo.SelectAll()
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("Product List:")
	for i, product := range products {
		fmt.Printf("%v .  - Category: %v\n     - Product Name: %v\n     - Price: %v\n", i+1, product.Category, product.ProductName, product.Price)
	}

	fmt.Printf("Select number to add cart: ")
	var selectProductNumber int
	fmt.Scan(&selectProductNumber)

	// Check if the selected product number is valid
	if selectProductNumber > 0 && selectProductNumber <= len(products) {
		product := products[selectProductNumber-1]
		// Call the function to add the selected product to the cart item list
		err := terminal.cartItemRepo.Add(product)
		if err != nil {
			return err
		}
	} else {
		fmt.Println("Invalid product number!")
		fmt.Printf("Choose the product again?(y/n): ")
		var tryAgain string
		fmt.Scan(&tryAgain)
		if tryAgain == "y" {
			// Back to User List Page
			terminal.ProductPage()
		}
		// Back to Dashboard Page
		terminal.DashboardPage(0)
	}

	// Back to Dashboard Page
	terminal.DashboardPage(0)
	return nil
}
