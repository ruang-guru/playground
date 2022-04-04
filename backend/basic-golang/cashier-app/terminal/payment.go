package terminal

import "fmt"

// PaymentPage is a function to enter the payment page and pay the bill
func (terminal *Terminal) PaymentPage() error {
	fmt.Printf("\x1bc") // clear screen

	fmt.Println()
	fmt.Println("Payment Page")
	sumPrice, _ := terminal.cartItemRepo.TotalPrice()
	fmt.Printf("Total price to pay: %v\n", sumPrice)
	fmt.Println()
	fmt.Printf("Enter your payment: ")
	var payment int
	fmt.Scan(&payment)

	// Return to dashboard page if payment is sufficient or more
	if payment >= sumPrice {
		terminal.DashboardPage(payment)
	} else {
		fmt.Println("Payment is not sufficient!")
		fmt.Printf("enter repayment?(y/n): ")
		var repayment string
		fmt.Scan(&repayment)
		if repayment == "y" {
			// Back to User List Page
			terminal.PaymentPage()
		}
		terminal.DashboardPage(0)
	}
	return nil
}
