package terminal

import (
	"fmt"
)

// UserListPage is a function to show user list
func (terminal *Terminal) UserListPage() error {
	fmt.Printf("\x1bc") // clear screen
	userList, err := terminal.usersRepo.SelectAll()
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("User List:")
	for i, user := range userList {
		fmt.Printf("%v . Username: %v, Status Login: %v\n", i+1, user.Username, user.Loggedin)
	}

	// Login
	terminal.LoginForm()
	return nil
}

// LoginForm is a function to enter the dashboard page
func (terminal *Terminal) LoginForm() error {
	var username, password, loginAgain string
	fmt.Println()
	fmt.Println("Login Form:")
	fmt.Printf("Username : ")
	fmt.Scan(&username)
	fmt.Printf("Password : ")
	fmt.Scan(&password)

	_, err := terminal.usersRepo.Login(username, password)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Printf("Loggin Again?(y/n): ")
		fmt.Scan(&loginAgain)
		if loginAgain == "y" {
			// Back to login form
			terminal.LoginForm()
		}
		fmt.Println("Thank you for using cashier application")
		return nil
	}

	// Show Dashboard Page after login success
	terminal.DashboardPage(0)
	return nil
}

// Logout is function to log out also change login status to false and reset cart
func (terminal *Terminal) Logout(username string) error {
	fmt.Printf("\x1bc") // clear screen
	err := terminal.usersRepo.Logout(username)
	if err != nil {
		return err
	}

	terminal.cartItemRepo.ResetCartItems() // reset cart items
	fmt.Printf("User %v is Logged out\n", username)
	return nil
}
