package repository

type User struct {
	Username string
	Password string
	Loggedin bool
}

type Product struct {
	Category    string
	ProductName string
	Price       int
}

type CartItem struct {
	Category    string
	ProductName string
	Price       int
	Quantity    int
}
