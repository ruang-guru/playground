package repository_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ruang-guru/playground/backend/basic-golang/cashier-app/db"
	repository "github.com/ruang-guru/playground/backend/basic-golang/cashier-app/repository"
)

var _ = Describe("Cashier Apps", func() {
	var testDB db.DB
	var userRepo repository.UserRepository
	var cartItemRepo repository.CartItemRepository
	var productRepo repository.ProductRepository
	var transactionRepo repository.TransactionRepository

	BeforeEach(func() {
		tables := map[db.TableName]db.Rows{
			"users": {
				db.Row{"username", "password", "loggedin"},
				db.Row{"aditira", "1234", "false"},
				db.Row{"dina", "4321", "false"},
				db.Row{"dito", "2552", "false"},
			},
			"products": {
				db.Row{"category", "product_name", "price"},
				db.Row{"Fruits", "Orange", "5000"},
				db.Row{"Fruits", "Apple", "2000"},
				db.Row{"Fruits", "Melon", "4000"},
				db.Row{"Fruits", "Watermelon", "10000"},
				db.Row{"Fruits", "Banana", "4000"},
				db.Row{"Vegetables", "Carrot", "2000"},
				db.Row{"Vegetables", "Broccoli", "5200"},
				db.Row{"Vegetables", "Cucumber", "3400"},
				db.Row{"Vegetables", "Potato", "6500"},
				db.Row{"Vegetables", "Tomato", "2200"},
				db.Row{"Drink", "Coffee", "4300"},
				db.Row{"Drink", "Milk", "4000"},
				db.Row{"Drink", "Tea", "2700"},
			},
			"cart_items": {
				db.Row{"category", "product_name", "price", "quantity"},
				db.Row{"Fruits", "Orange", "5000", "18"},
				db.Row{"Fruits", "Apple", "2000", "18"},
				db.Row{"Fruits", "Melon", "4000", "12"},
			},
		}
		testDB = db.NewMemoryDB(tables)

		userRepo = repository.NewUserRepository(testDB)
		cartItemRepo = repository.NewCartItemRepository(testDB)
		productRepo = repository.NewProductRepository(testDB)
		transactionRepo = repository.NewTransactionRepository(cartItemRepo)
	})

	Describe("Select All Users", func() {
		When("get all user list from file user.csv", func() {
			It("returns user information in the same order as user.csv", func() {
				userList, err := userRepo.SelectAll()
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
			It("make other users logged out", func() {
				_, err := userRepo.Login("aditira", "1234")
				Expect(err).ToNot(HaveOccurred())
				loggedInUser, _ := userRepo.FindLoggedinUser()
				Expect(*loggedInUser).To(Equal("aditira"))

				_, err = userRepo.Login("dina", "4321")
				Expect(err).ToNot(HaveOccurred())
				loggedInUser, _ = userRepo.FindLoggedinUser()
				Expect(*loggedInUser).To(Equal("dina"))
			})
			It("change login status to true", func() {
				_, err := userRepo.Login("aditira", "1234")
				Expect(err).ToNot(HaveOccurred())
				loggedInUser, _ := userRepo.FindLoggedinUser()
				Expect(*loggedInUser).To(Equal("aditira"))
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

	Describe("Logout", func() {
		When("the user is logged in", func() {
			It("accepts the logout", func() {
				userRepo.Login("aditira", "1234")
				err := userRepo.Logout("aditira")
				Expect(err).To(BeNil())
			})
		})
		When("the user is not logged in", func() {
			It("rejects the logout", func() {
				userRepo.LogoutAll()
				err := userRepo.Logout("aditira")
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("Find logged in user", func() {
		When("there is logged in user", func() {
			It("return user name", func() {
				userRepo.Login("aditira", "1234")
				username, _ := userRepo.FindLoggedinUser()
				Expect(*username).To(Equal("aditira"))
			})
		})
		When("there is not logged in user", func() {
			It("return message `No User Logged In!`", func() {
				userRepo.LogoutAll()
				_, err := userRepo.FindLoggedinUser()
				Expect(err.Error()).To(Equal("no user is logged in"))
			})
		})
	})

	Describe("Select All Product", func() {
		When("get all product list from file product.csv", func() {
			It("returns product information in the same order as product.csv", func() {
				productList, err := productRepo.SelectAll()
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
			It("adds the item to the cart", func() {
				cartItemRepo.ResetCartItems()
				product, err := productRepo.SelectAll()
				Expect(err).ToNot(HaveOccurred())
				cartItemRepo.Add(product[0])
				cartItemRepo.Add(product[1])
				cartItemRepo.Add(product[1])
				cartItem, err := cartItemRepo.SelectAll()
				Expect(err).ToNot(HaveOccurred())
				Expect(cartItem[0].ProductName).To(Equal("Orange"))
				Expect(cartItem[0].Quantity).To(Equal(1))
				Expect(cartItem[0].Price).To(Equal(5000))
				Expect(cartItem[1].ProductName).To(Equal("Apple"))
				Expect(cartItem[1].Quantity).To(Equal(2))
				Expect(cartItem[1].Price).To(Equal(2000))
			})
		})
	})

	Describe("Select All Cart Items", func() {
		When("there is cart item", func() {
			It("returns cart items information in the same order as cart_items.csv", func() {
				cartItemRepo.ResetCartItems()
				product, err := productRepo.SelectAll()
				Expect(err).ToNot(HaveOccurred())
				cartItemRepo.Add(product[0])
				cartItemRepo.Add(product[1])
				cartItemRepo.Add(product[1])
				cartItem, err := cartItemRepo.SelectAll()
				Expect(err).ToNot(HaveOccurred())
				Expect(cartItem[0].ProductName).To(Equal("Orange"))
				Expect(cartItem[0].Quantity).To(Equal(1))
				Expect(cartItem[0].Price).To(Equal(5000))
				Expect(cartItem[1].ProductName).To(Equal("Apple"))
				Expect(cartItem[1].Quantity).To(Equal(2))
				Expect(cartItem[1].Price).To(Equal(2000))
			})
		})
	})

	Describe("Total Price", func() {
		When("there is cart item", func() {
			It("returns total price", func() {
				cartItemRepo.ResetCartItems()
				product, err := productRepo.SelectAll()
				Expect(err).ToNot(HaveOccurred())
				cartItemRepo.Add(product[0])
				cartItemRepo.Add(product[1])
				cartItemRepo.Add(product[1])
				totalPrice, err := cartItemRepo.TotalPrice()
				Expect(err).ToNot(HaveOccurred())
				Expect(totalPrice).To(Equal(9000))
			})
		})
	})

	Describe("Reset Cart Items", func() {
		When("there is cart item", func() {
			It("returns empty cart items", func() {
				product, err := productRepo.SelectAll()
				Expect(err).ToNot(HaveOccurred())
				cartItemRepo.Add(product[0])
				cartItemRepo.Add(product[1])
				cartItemRepo.Add(product[1])
				cartItemRepo.ResetCartItems()
				cartItem, err := cartItemRepo.SelectAll()
				Expect(err).ToNot(HaveOccurred())
				Expect(cartItem).To(BeEmpty())
			})
		})
	})

	Describe("Pay", func() {
		When("pay the total price ", func() {
			It("will return the money changes", func() {
				cartItemRepo.ResetCartItems()
				product, err := productRepo.SelectAll()
				Expect(err).ToNot(HaveOccurred())
				cartItemRepo.Add(product[0])
				cartItemRepo.Add(product[1])
				cartItemRepo.Add(product[1])
				Expect(transactionRepo.Pay(50000)).To(Equal(41000))
			})
		})
	})

})
