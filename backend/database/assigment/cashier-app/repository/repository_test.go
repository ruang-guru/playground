package repository_test

import (
	"database/sql"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ruang-guru/playground/backend/database/assigment/cashier-app/repository"
)

var _ = Describe("Repository Test", func() {
	var db *sql.DB
	var err error
	var userRepo *repository.UserRepository
	var cartItemRepo *repository.CartItemRepository
	var productRepo *repository.ProductRepository
	var transactionRepo repository.TransactionRepository
	var salesRepo *repository.SalesRepository

	BeforeEach(func() {
		// Setup
		db, err = sql.Open("sqlite3", "./cashier-app.db")
		if err != nil {
			panic(err)
		}

		_, err = db.Exec(`
		CREATE TABLE users (
			id integer not null primary key AUTOINCREMENT,
			username varchar(255) not null,
			password varchar(255) not null,
			role varchar(255) not null,
			loggedin boolean not null
		);

		CREATE TABLE products (
			id integer not null primary key AUTOINCREMENT,
			product_name varchar(255) not null,
			category varchar(255) not null,
			price integer not null
		);
		
		CREATE TABLE cart_items (
			id integer not null primary key AUTOINCREMENT,
			product_id integer not null,
			quantity integer not null,
			FOREIGN KEY (product_id) REFERENCES products(id)
		);
		
		CREATE TABLE sales (
			id integer not null primary key AUTOINCREMENT,
			date DATE not null,
			product_id integer not null,
			quantity integer not null,
			FOREIGN KEY (product_id) REFERENCES products(id)
		);
		
		INSERT INTO users(username, password, role, loggedin) VALUES
			('aditira', '1234', 'admin', false),
			('dina', '4321', 'employee', false),
			('dito', '2552', 'employee', false);
		
		INSERT INTO products(product_name, category, price) VALUES
			('Orange', 'Fruits', 5000),
			('Apple', 'Fruits', 2000),
			('Melon', 'Fruits', 4000),
			('Watermelon', 'Fruits', 10000),
			('Banana', 'Fruits', 4000),
			('Carrot', 'Vegetables', 2000),
			('Broccoli', 'Vegetables', 5200),
			('Cucumber', 'Vegetables', 3400),
			('Potato', 'Vegetables', 6500),
			('Tomato', 'Vegetables', 2200),
			('Coffee', 'Drink', 4300),
			('Milk', 'Drink', 4000),
			('Tea', 'Drink', 2700);`)

		if err != nil {
			panic(err)
		}

		userRepo = repository.NewUserRepository(db)
		productRepo = repository.NewProductRepository(db)
		cartItemRepo = repository.NewCartItemRepository(db)
		salesRepo = repository.NewSalesRepository(db)
		transactionRepo = repository.NewTransactionRepository(*cartItemRepo, *salesRepo)
	})

	AfterEach(func() {
		// Teardown
		db, err := sql.Open("sqlite3", "./cashier-app.db")
		if err != nil {
			panic(err)
		}

		_, err = db.Exec(`
		DROP TABLE users;
		DROP TABLE products;
		DROP TABLE cart_items;
		DROP TABLE sales;`)

		if err != nil {
			panic(err)
		}

		os.Remove("./cashier-app.db")
	})

	Describe("Select All Users", func() {
		When("get all user list from database", func() {
			It("should return all user list", func() {
				userList, err := userRepo.FetchUsers()
				Expect(err).ToNot(HaveOccurred())

				Expect(userList[0].Username).To(Equal("aditira"))
				Expect(userList[0].Password).To(Equal("1234"))
				Expect(userList[0].Loggedin).To(Equal(false))
				Expect(userList[1].Username).To(Equal("dina"))
				Expect(userList[1].Password).To(Equal("4321"))
				Expect(userList[1].Loggedin).To(Equal(false))
				Expect(userList[2].Username).To(Equal("dito"))
				Expect(userList[2].Password).To(Equal("2552"))
				Expect(userList[2].Loggedin).To(Equal(false))
			})
		})
	})

	Describe("Login", func() {
		When("username and password are correct", func() {
			It("accepts the login", func() {
				res, err := userRepo.Login("aditira", "1234")
				Expect(err).ToNot(HaveOccurred())
				Expect(*res).To(Equal("aditira"))
			})
		})
		When("username is correct but password is incorrect", func() {
			It("rejects the login", func() {
				_, err := userRepo.Login("aditira", "4567")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("Login Failed"))
			})
		})
		When("both username and password is incorrect", func() {
			It("rejects the login", func() {
				_, err := userRepo.Login("juno", "23885")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("Login Failed"))
			})
		})
	})

	Describe("Select All Product", func() {
		When("get all product list from database in products table", func() {
			It("returns product information in the same order as database", func() {
				productList, err := productRepo.FetchProducts()
				Expect(err).ToNot(HaveOccurred())
				Expect(productList[0].Category).To(Equal("Fruits"))
				Expect(productList[0].ProductName).To(Equal("Orange"))
				Expect(productList[0].Price).To(Equal(5000))
				Expect(productList[1].Category).To(Equal("Fruits"))
				Expect(productList[1].ProductName).To(Equal("Apple"))
				Expect(productList[1].Price).To(Equal(2000))
			})
		})
	})

	Describe("Add Cart Item", func() {
		When("select multiple items and confirm will add items to cart", func() {
			It("add the item to the cart", func() {
				cartItemRepo.ResetCartItems()
				product, err := productRepo.FetchProductByName("Orange")
				Expect(err).ToNot(HaveOccurred())
				_, err = cartItemRepo.FetchCartByProductID(product.ID)
				Expect(err).To(HaveOccurred())
				cartItemRepo.InsertCartItem(repository.CartItem{
					ProductID: product.ID,
					Quantity:  1,
				})
				cartItem, err := cartItemRepo.FetchCartItems()
				Expect(err).ToNot(HaveOccurred())
				Expect(cartItem[0].ProductName).To(Equal("Orange"))
				Expect(cartItem[0].Quantity).To(Equal(1))
				Expect(cartItem[0].Price).To(Equal(5000))
			})
			It("update item to product cart", func() {
				cartItemRepo.ResetCartItems()
				product, err := productRepo.FetchProductByName("Orange")
				Expect(err).ToNot(HaveOccurred())
				_, err = cartItemRepo.FetchCartByProductID(product.ID)
				Expect(err).To(HaveOccurred())
				cartItemRepo.InsertCartItem(repository.CartItem{
					ProductID: product.ID,
					Quantity:  1,
				})
				cartItem, err := cartItemRepo.FetchCartByProductID(product.ID)
				Expect(err).ToNot(HaveOccurred())
				err = cartItemRepo.UpdateCartItemQuantity(cartItem)
				Expect(err).ToNot(HaveOccurred())
				cartItems, err := cartItemRepo.FetchCartItems()
				Expect(err).ToNot(HaveOccurred())
				Expect(cartItems[0].ProductName).To(Equal("Orange"))
				Expect(cartItems[0].Quantity).To(Equal(2))
				Expect(cartItems[0].Price).To(Equal(5000))
			})
		})
	})

	Describe("Total Price", func() {
		When("there is cart item", func() {
			It("returns total price", func() {
				cartItemRepo.ResetCartItems()
				product, err := productRepo.FetchProductByName("Orange")
				Expect(err).ToNot(HaveOccurred())
				_, err = cartItemRepo.FetchCartByProductID(product.ID)
				Expect(err).To(HaveOccurred())
				cartItemRepo.InsertCartItem(repository.CartItem{
					ProductID: product.ID,
					Quantity:  1,
				})
				product, err = productRepo.FetchProductByName("Apple")
				Expect(err).ToNot(HaveOccurred())
				_, err = cartItemRepo.FetchCartByProductID(product.ID)
				Expect(err).To(HaveOccurred())
				cartItemRepo.InsertCartItem(repository.CartItem{
					ProductID: product.ID,
					Quantity:  1,
				})
				totalPrice, err := cartItemRepo.TotalPrice()
				Expect(err).ToNot(HaveOccurred())
				Expect(totalPrice).To(Equal(7000))
			})
		})
	})

	Describe("Reset Cart Items", func() {
		When("there is cart item", func() {
			It("returns empty cart items", func() {
				cartItemRepo.ResetCartItems()
				cartItem, err := cartItemRepo.FetchCartItems()
				Expect(err).ToNot(HaveOccurred())
				Expect(cartItem).To(BeEmpty())
			})
		})
	})

	Describe("Pay", func() {
		When("pay the total price ", func() {
			It("will return the money changes", func() {
				cartItemRepo.ResetCartItems()
				product, err := productRepo.FetchProductByName("Orange")
				Expect(err).ToNot(HaveOccurred())
				_, err = cartItemRepo.FetchCartByProductID(product.ID)
				Expect(err).To(HaveOccurred())
				cartItemRepo.InsertCartItem(repository.CartItem{
					ProductID: product.ID,
					Quantity:  1,
				})
				product, err = productRepo.FetchProductByName("Apple")
				Expect(err).ToNot(HaveOccurred())
				_, err = cartItemRepo.FetchCartByProductID(product.ID)
				Expect(err).To(HaveOccurred())
				cartItemRepo.InsertCartItem(repository.CartItem{
					ProductID: product.ID,
					Quantity:  1,
				})
				Expect(transactionRepo.Pay([]repository.CartItem{}, 50000)).To(Equal(43000))
			})
		})
	})
})
