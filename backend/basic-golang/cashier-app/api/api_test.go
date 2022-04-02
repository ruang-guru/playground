package api_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/ruang-guru/playground/backend/basic-golang/cashier-app/api"
	"github.com/ruang-guru/playground/backend/basic-golang/cashier-app/db"
	"github.com/ruang-guru/playground/backend/basic-golang/cashier-app/repository"
)

var _ = Describe("Api", func() {
	var testDB db.DB
	var userRepo repository.UserRepository
	var cartItemRepo repository.CartItemRepository
	var productRepo repository.ProductRepository
	var transactionRepo repository.TransactionRepository
	var server api.API

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
				db.Row{"Vegetables", "Carrot", "2000"},
				db.Row{"Vegetables", "Tomato", "2200"},
				db.Row{"Drink", "Tea", "2700"},
			},
			"cart_items": {
				db.Row{"category", "product_name", "price", "quantity"},
			},
		}
		testDB = db.NewMemoryDB(tables)

		userRepo = repository.NewUserRepository(testDB)
		cartItemRepo = repository.NewCartItemRepository(testDB)
		productRepo = repository.NewProductRepository(testDB)
		transactionRepo = repository.NewTransactionRepository(cartItemRepo)

		server = api.NewAPI(userRepo, productRepo, cartItemRepo, transactionRepo)
	})

	Describe("/user/login", func() {
		When("the username and password are correct", func() {
			It("should return a successful login response", func() {
				wr := httptest.NewRecorder()
				req := httptest.NewRequest(http.MethodGet, "/api/user/login?username=aditira&password=1234", nil)
				server.Handler().ServeHTTP(wr, req)

				Expect(wr.Result().StatusCode).To(Equal(http.StatusOK))

				loginSuccessResponse := api.LoginSuccessResponse{}
				json.NewDecoder(wr.Body).Decode(&loginSuccessResponse)
				Expect(loginSuccessResponse.Username).To(Equal("aditira"))
			})
		})
		When("the username and password are incorrect", func() {
			It("should return unauthorized", func() {
				wr := httptest.NewRecorder()
				req := httptest.NewRequest(http.MethodGet, "/api/user/login?username=aditira&password=123", nil)
				server.Handler().ServeHTTP(wr, req)

				Expect(wr.Result().StatusCode).To(Equal(http.StatusUnauthorized))
			})
		})
	})

	Describe("/user/logout", func() {
		When("the username is logged in", func() {
			It("should return a successful logout response", func() {
				//login
				wr := httptest.NewRecorder()
				req := httptest.NewRequest(http.MethodGet, "/api/user/login?username=aditira&password=1234", nil)
				server.Handler().ServeHTTP(wr, req)

				//logout
				wr = httptest.NewRecorder()
				req = httptest.NewRequest(http.MethodGet, "/api/user/logout?username=aditira", nil)
				server.Handler().ServeHTTP(wr, req)

				Expect(wr.Result().StatusCode).To(Equal(http.StatusOK))
			})
		})
		When("the username is not logged in", func() {
			It("should return unauthorized", func() {
				wr := httptest.NewRecorder()
				req := httptest.NewRequest(http.MethodGet, "/api/user/logout?username=aditira", nil)
				server.Handler().ServeHTTP(wr, req)

				Expect(wr.Result().StatusCode).To(Equal(http.StatusUnauthorized))
			})
		})
	})

	Describe("cart/add", func() {
		When("the product is found", func() {
			It("returns the added product detail", func() {
				//login
				wr := httptest.NewRecorder()
				req := httptest.NewRequest(http.MethodGet, "/api/user/login?username=aditira&password=1234", nil)
				server.Handler().ServeHTTP(wr, req)

				//add to cart
				wr = httptest.NewRecorder()
				req = httptest.NewRequest(http.MethodGet, "/api/cart/add?product_name=Tomato", nil)
				server.Handler().ServeHTTP(wr, req)

				Expect(wr.Result().StatusCode).To(Equal(http.StatusOK))

				addToCartSuccessResponse := api.AddToCartSuccessResponse{}
				json.NewDecoder(wr.Body).Decode(&addToCartSuccessResponse)

				Expect(addToCartSuccessResponse.Name).To(Equal("Tomato"))
				Expect(addToCartSuccessResponse.Category).To(Equal("Vegetables"))
				Expect(addToCartSuccessResponse.Price).To(Equal(2200))
			})

			It("adds the cart items", func() {
				//login
				wr := httptest.NewRecorder()
				req := httptest.NewRequest(http.MethodGet, "/api/user/login?username=aditira&password=1234", nil)
				server.Handler().ServeHTTP(wr, req)

				//add to cart
				wr = httptest.NewRecorder()
				req = httptest.NewRequest(http.MethodGet, "/api/cart/add?product_name=Tomato", nil)
				server.Handler().ServeHTTP(wr, req)

				//get cart items
				wr = httptest.NewRecorder()
				req = httptest.NewRequest(http.MethodGet, "/api/carts", nil)
				server.Handler().ServeHTTP(wr, req)

				cartListSuccessResponse := api.CartListSuccessResponse{}
				json.NewDecoder(wr.Body).Decode(&cartListSuccessResponse)

				Expect(cartListSuccessResponse.CartItems).To(HaveLen(1))
				Expect(cartListSuccessResponse.CartItems[0].ProductName).To(Equal("Tomato"))
				Expect(cartListSuccessResponse.CartItems[0].Quantity).To(Equal(1))
				Expect(cartListSuccessResponse.CartItems[0].Price).To(Equal(2200))
			})

			When("there are multiple similar product", func() {
				It("appends the quantity", func() {
					//login
					wr := httptest.NewRecorder()
					req := httptest.NewRequest(http.MethodGet, "/api/user/login?username=aditira&password=1234", nil)
					server.Handler().ServeHTTP(wr, req)

					//add to cart
					wr = httptest.NewRecorder()
					req = httptest.NewRequest(http.MethodGet, "/api/cart/add?product_name=Tomato", nil)
					server.Handler().ServeHTTP(wr, req)

					//add to cart again
					wr = httptest.NewRecorder()
					req = httptest.NewRequest(http.MethodGet, "/api/cart/add?product_name=Tomato", nil)
					server.Handler().ServeHTTP(wr, req)

					//get cart items
					wr = httptest.NewRecorder()
					req = httptest.NewRequest(http.MethodGet, "/api/carts", nil)
					server.Handler().ServeHTTP(wr, req)

					cartListSuccessResponse := api.CartListSuccessResponse{}
					json.NewDecoder(wr.Body).Decode(&cartListSuccessResponse)

					Expect(cartListSuccessResponse.CartItems).To(HaveLen(1))
					Expect(cartListSuccessResponse.CartItems[0].ProductName).To(Equal("Tomato"))
					Expect(cartListSuccessResponse.CartItems[0].Quantity).To(Equal(2))
					Expect(cartListSuccessResponse.CartItems[0].Price).To(Equal(2200))
				})
			})

			When("there are multiple different products", func() {
				It("returns all those products with quantity = 1", func() {
					//login
					wr := httptest.NewRecorder()
					req := httptest.NewRequest(http.MethodGet, "/api/user/login?username=aditira&password=1234", nil)
					server.Handler().ServeHTTP(wr, req)

					//add to cart
					wr = httptest.NewRecorder()
					req = httptest.NewRequest(http.MethodGet, "/api/cart/add?product_name=Tomato", nil)
					server.Handler().ServeHTTP(wr, req)

					//add to cart again
					wr = httptest.NewRecorder()
					req = httptest.NewRequest(http.MethodGet, "/api/cart/add?product_name=Tea", nil)
					server.Handler().ServeHTTP(wr, req)

					//get cart items
					wr = httptest.NewRecorder()
					req = httptest.NewRequest(http.MethodGet, "/api/carts", nil)
					server.Handler().ServeHTTP(wr, req)

					cartListSuccessResponse := api.CartListSuccessResponse{}
					json.NewDecoder(wr.Body).Decode(&cartListSuccessResponse)

					Expect(cartListSuccessResponse.CartItems).To(HaveLen(2))

					Expect(cartListSuccessResponse.CartItems[0].ProductName).To(Equal("Tomato"))
					Expect(cartListSuccessResponse.CartItems[0].Quantity).To(Equal(1))
					Expect(cartListSuccessResponse.CartItems[0].Price).To(Equal(2200))

					Expect(cartListSuccessResponse.CartItems[1].ProductName).To(Equal("Tea"))
					Expect(cartListSuccessResponse.CartItems[1].Quantity).To(Equal(1))
					Expect(cartListSuccessResponse.CartItems[1].Price).To(Equal(2700))
				})
			})
		})
		When("the product is not found", func() {
			It("returns not found (404)", func() {
				//login
				wr := httptest.NewRecorder()
				req := httptest.NewRequest(http.MethodGet, "/api/user/login?username=aditira&password=1234", nil)
				server.Handler().ServeHTTP(wr, req)

				//add to cart
				wr = httptest.NewRecorder()
				req = httptest.NewRequest(http.MethodGet, "/api/cart/add?product_name=Chocolate", nil)
				server.Handler().ServeHTTP(wr, req)

				Expect(wr.Code).To(Equal(http.StatusNotFound))
			})
		})
	})

	Describe("cart/clear", func() {
		It("clears the cart items", func() {
			//login
			wr := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/api/user/login?username=aditira&password=1234", nil)
			server.Handler().ServeHTTP(wr, req)

			//add to cart
			wr = httptest.NewRecorder()
			req = httptest.NewRequest(http.MethodGet, "/api/cart/add?product_name=Tomato", nil)
			server.Handler().ServeHTTP(wr, req)

			//clear cart
			wr = httptest.NewRecorder()
			req = httptest.NewRequest(http.MethodGet, "/api/cart/clear", nil)
			server.Handler().ServeHTTP(wr, req)

			//get cart items
			wr = httptest.NewRecorder()
			req = httptest.NewRequest(http.MethodGet, "/api/carts", nil)
			server.Handler().ServeHTTP(wr, req)

			cartListSuccessResponse := api.CartListSuccessResponse{}
			json.NewDecoder(wr.Body).Decode(&cartListSuccessResponse)
			Expect(cartListSuccessResponse.CartItems).To(HaveLen(0))
		})
	})

	Describe("products", func() {
		It("returns all products", func() {
			wr := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/api/products", nil)
			server.Handler().ServeHTTP(wr, req)

			productListSuccessResponse := api.ProductListSuccessResponse{}
			json.NewDecoder(wr.Body).Decode(&productListSuccessResponse)

			Expect(productListSuccessResponse.Products).To(HaveLen(4))
			Expect(productListSuccessResponse.Products[0].Name).To(Equal("Orange"))
			Expect(productListSuccessResponse.Products[0].Price).To(Equal(5000))
			Expect(productListSuccessResponse.Products[0].Category).To(Equal("Fruits"))
		})
	})

	Describe("dashboard", func() {
		When("username is not logged in", func() {
			It("return unauthorized", func() {
				wr := httptest.NewRecorder()
				req := httptest.NewRequest(http.MethodGet, "/api/dashboard", nil)
				server.Handler().ServeHTTP(wr, req)

				Expect(wr.Code).To(Equal(http.StatusUnauthorized))
			})
		})
		When("username is logged in", func() {
			It("returns the username", func() {
				//login
				wr := httptest.NewRecorder()
				req := httptest.NewRequest(http.MethodGet, "/api/user/login?username=aditira&password=1234", nil)
				server.Handler().ServeHTTP(wr, req)

				//dashboard
				wr = httptest.NewRecorder()
				req = httptest.NewRequest(http.MethodGet, "/api/dashboard", nil)
				server.Handler().ServeHTTP(wr, req)

				dashboardSuccessResponse := api.DashboardSuccessResponse{}
				json.NewDecoder(wr.Body).Decode(&dashboardSuccessResponse)

				Expect(dashboardSuccessResponse.Username).To(Equal("aditira"))
			})

			It("returns the cart items", func() {
				//login
				wr := httptest.NewRecorder()
				req := httptest.NewRequest(http.MethodGet, "/api/user/login?username=aditira&password=1234", nil)
				server.Handler().ServeHTTP(wr, req)

				//add to cart
				wr = httptest.NewRecorder()
				req = httptest.NewRequest(http.MethodGet, "/api/cart/add?product_name=Tomato", nil)
				server.Handler().ServeHTTP(wr, req)

				//dashboard
				wr = httptest.NewRecorder()
				req = httptest.NewRequest(http.MethodGet, "/api/dashboard", nil)
				server.Handler().ServeHTTP(wr, req)

				dashboardSuccessResponse := api.DashboardSuccessResponse{}
				json.NewDecoder(wr.Body).Decode(&dashboardSuccessResponse)

				Expect(dashboardSuccessResponse.CartItems).To(HaveLen(1))
				Expect(dashboardSuccessResponse.CartItems[0].ProductName).To(Equal("Tomato"))
				Expect(dashboardSuccessResponse.CartItems[0].Price).To(Equal(2200))
				Expect(dashboardSuccessResponse.CartItems[0].Quantity).To(Equal(1))
			})

			It("returns the total price", func() {
				//login
				wr := httptest.NewRecorder()
				req := httptest.NewRequest(http.MethodGet, "/api/user/login?username=aditira&password=1234", nil)
				server.Handler().ServeHTTP(wr, req)

				//add to cart
				wr = httptest.NewRecorder()
				req = httptest.NewRequest(http.MethodGet, "/api/cart/add?product_name=Tomato", nil)
				server.Handler().ServeHTTP(wr, req)

				//add to cart again
				wr = httptest.NewRecorder()
				req = httptest.NewRequest(http.MethodGet, "/api/cart/add?product_name=Tea", nil)
				server.Handler().ServeHTTP(wr, req)

				//dashboard
				wr = httptest.NewRecorder()
				req = httptest.NewRequest(http.MethodGet, "/api/dashboard", nil)
				server.Handler().ServeHTTP(wr, req)

				dashboardSuccessResponse := api.DashboardSuccessResponse{}
				json.NewDecoder(wr.Body).Decode(&dashboardSuccessResponse)

				Expect(dashboardSuccessResponse.TotalPrice).To(Equal(4900))
			})

			It("calculate the change", func() {
				//login
				wr := httptest.NewRecorder()
				req := httptest.NewRequest(http.MethodGet, "/api/user/login?username=aditira&password=1234", nil)
				server.Handler().ServeHTTP(wr, req)

				//add to cart
				wr = httptest.NewRecorder()
				req = httptest.NewRequest(http.MethodGet, "/api/cart/add?product_name=Tomato", nil)
				server.Handler().ServeHTTP(wr, req)

				//dashboard
				wr = httptest.NewRecorder()
				req = httptest.NewRequest(http.MethodGet, "/api/dashboard?cash=4900", nil)
				server.Handler().ServeHTTP(wr, req)

				dashboardSuccessResponse := api.DashboardSuccessResponse{}
				json.NewDecoder(wr.Body).Decode(&dashboardSuccessResponse)

				Expect(dashboardSuccessResponse.TotalPrice).To(Equal(2200))
				Expect(dashboardSuccessResponse.MoneyChanges).To(Equal(2700))
			})
		})
	})
})
