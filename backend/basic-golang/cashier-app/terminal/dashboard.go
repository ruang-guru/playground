package terminal

import "fmt"

// DashboardPage is the main page of the cashier app
func (terminal *Terminal) DashboardPage(cash int) error {
	fmt.Printf("\x1bc") // clear screen

	// Find user with loggedin status is true
	username, err := terminal.usersRepo.FindLoggedinUser()
	if err != nil {
		return err
	}

	// Pay is the function to pay the bill
	moneyChanges, err := terminal.transactionRepo.Pay(cash)
	if err != nil {
		return err
	}

	// Select all cart items
	cartItems, err := terminal.cartItemRepo.SelectAll()
	if err != nil {
		return err
	}

	// Find the total price of all cart items
	sumPrice, err := terminal.cartItemRepo.TotalPrice()
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("Dashboard Page")
	fmt.Printf("- Logged in with username: %v\n", *username)
	fmt.Printf("- Money Changes: %v\n", moneyChanges)
	fmt.Println("- Cart Items:")
	if len(cartItems) > 0 {
		for i, cartItem := range cartItems {
			fmt.Printf("  %v. Product Name: %v, Price: %v, Quantity: %v\n", i+1, cartItem.ProductName, cartItem.Price, cartItem.Quantity)
		}
		fmt.Printf("- Total Price: %v\n", sumPrice)
	} else {
		fmt.Println("  0. Cart is empty!")
	}

	fmt.Println()
	fmt.Println("Main menu:")
	fmt.Println("1. Order")
	fmt.Println("2. Pay")
	fmt.Println("3. Logout")
	fmt.Printf("Please choose menu: ")
	var choiceMenu int
	fmt.Scan(&choiceMenu)

	// Call the function according to the selected main menu
	switch choiceMenu {
	case 1:
		// Call the product page function to show the product list and add new product to the cart item list
		terminal.ProductPage()
	case 2:
		// Call the payment page function to enter the payment
		terminal.PaymentPage()
	case 3:
		// Call the logout function to logout the user
		terminal.Logout(*username)
	default:
		fmt.Println("Invalid choice!")
		// Call the dashboard page function to show the dashboard page again
		terminal.DashboardPage(0)
	}

	return nil
}
