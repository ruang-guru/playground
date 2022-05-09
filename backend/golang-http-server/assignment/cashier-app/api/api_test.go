package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/cashier-app/api"
	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/cashier-app/db"
	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/cashier-app/repository"
)

var _ = Describe("Api", func() {
	var testDB db.DB
	var userRepo repository.UserRepository
	var cartItemRepo repository.CartItemRepository
	var productRepo repository.ProductRepository
	var transactionRepo repository.TransactionRepository
	var salesRepo repository.SalesRepository
	var server api.API

	BeforeEach(func() {
		tables := map[db.TableName]db.Rows{
			"users": {
				db.Row{"username", "password", "loggedin", "role"},
				db.Row{"aditira", "1234", "false", "admin"},
				db.Row{"dina", "4321", "false", "employee"},
				db.Row{"dito", "2552", "false", "employee"},
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
			"sales": {
				db.Row{"date", "category", "product", "quantity", "price", "total"},
				db.Row{"2022-05-03 22:10:26", "Fruits", "Orange", "1", "5000", "5000"},
				db.Row{"2022-05-01 22:10:26", "Fruits", "Apple", "1", "2000", "2000"},
				db.Row{"2022-05-01 22:10:26", "Fruits", "Orange", "1", "5000", "5000"},
				db.Row{"2022-05-01 22:10:26", "Fruits", "Apple", "1", "2000", "2000"},
			},
		}
		testDB = db.NewMemoryDB(tables)

		userRepo = repository.NewUserRepository(testDB)
		cartItemRepo = repository.NewCartItemRepository(testDB)
		productRepo = repository.NewProductRepository(testDB)
		salesRepo = repository.NewSalesRepository(testDB)
		transactionRepo = repository.NewTransactionRepository(cartItemRepo, salesRepo)

		server = api.NewAPI(userRepo, productRepo, cartItemRepo, transactionRepo, salesRepo)
	})

	Describe("/user/login", func() {
		When("the username and password are correct", func() {
			It("should return a successful login response", func() {
				var jsonStr = []byte(`{"username": "aditira", "password": "1234"}`)
				wr := httptest.NewRecorder()
				req := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(jsonStr))
				server.Handler().ServeHTTP(wr, req)

				Expect(wr.Result().StatusCode).To(Equal(http.StatusOK))

				loginSuccessResponse := api.LoginSuccessResponse{}
				json.NewDecoder(wr.Body).Decode(&loginSuccessResponse)
				Expect(loginSuccessResponse.Username).To(Equal("aditira"))
			})
		})
		When("the username and password are incorrect", func() {
			It("should return unauthorized", func() {
				var jsonStr = []byte(`{"username": "aditira", "password": "12345"}`)
				wr := httptest.NewRecorder()
				req := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(jsonStr))
				server.Handler().ServeHTTP(wr, req)

				Expect(wr.Result().StatusCode).To(Equal(http.StatusUnauthorized))
			})
		})
	})

	Describe("/user/logout", func() {
		When("the username is logged in", func() {
			It("should return successful response", func() {
				//login
				var jsonStr = []byte(`{"username": "aditira", "password": "1234"}`)
				wr := httptest.NewRecorder()
				req := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(jsonStr))
				server.Handler().ServeHTTP(wr, req)

				var cookie *http.Cookie
				for _, c := range wr.Result().Cookies() {
					if c.Name == "token" {
						cookie = c
					}
				}

				//logout
				wr = httptest.NewRecorder()
				req = httptest.NewRequest(http.MethodPost, "/api/user/logout", nil)
				req.AddCookie(cookie)
				server.Handler().ServeHTTP(wr, req)

				Expect(wr.Result().StatusCode).To(Equal(http.StatusOK))
			})
		})
		When("the username is not logged in", func() {
			It("should return unauthorized", func() {
				wr := httptest.NewRecorder()
				req := httptest.NewRequest(http.MethodPost, "/api/user/logout", nil)
				server.Handler().ServeHTTP(wr, req)

				Expect(wr.Result().StatusCode).To(Equal(http.StatusUnauthorized))
			})
		})
	})

	Describe("cart/add", func() {
		When("the product is found", func() {
			It("returns the added product detail", func() {
				//login
				var jsonStr = []byte(`{"username": "aditira", "password": "1234"}`)
				wr := httptest.NewRecorder()
				req := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(jsonStr))
				server.Handler().ServeHTTP(wr, req)

				Expect(wr.Code).To(Equal(200))

				var cookie *http.Cookie
				for _, c := range wr.Result().Cookies() {
					if c.Name == "token" {
						cookie = c
					}
				}

				//add to cart
				var productJsonStr = []byte(`{"product_name": "Tomato"}`)
				wr = httptest.NewRecorder()
				req = httptest.NewRequest(http.MethodPost, "/api/cart/add", bytes.NewBuffer(productJsonStr))
				req.AddCookie(cookie)
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
				var jsonStr = []byte(`{"username": "aditira", "password": "1234"}`)
				wr := httptest.NewRecorder()
				req := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(jsonStr))
				server.Handler().ServeHTTP(wr, req)

				Expect(wr.Code).To(Equal(200))

				var cookie *http.Cookie
				for _, c := range wr.Result().Cookies() {
					if c.Name == "token" {
						cookie = c
					}
				}

				//add to cart
				var productJsonStr []byte
				productJsonStr = []byte(`{"product_name": "Tomato"}`)
				wr = httptest.NewRecorder()
				req = httptest.NewRequest(http.MethodPost, "/api/cart/add", bytes.NewBuffer(productJsonStr))
				req.AddCookie(cookie)
				server.Handler().ServeHTTP(wr, req)

				Expect(wr.Code).To(Equal(200))

				//get cart items
				wr = httptest.NewRecorder()
				req = httptest.NewRequest(http.MethodGet, "/api/carts", nil)
				req.AddCookie(cookie)
				server.Handler().ServeHTTP(wr, req)

				Expect(wr.Code).To(Equal(200))

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
					var jsonStr = []byte(`{"username": "aditira", "password": "1234"}`)
					wr := httptest.NewRecorder()
					req := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(jsonStr))
					server.Handler().ServeHTTP(wr, req)

					Expect(wr.Code).To(Equal(200))

					var cookie *http.Cookie
					for _, c := range wr.Result().Cookies() {
						if c.Name == "token" {
							cookie = c
						}
					}

					//add to cart
					var productJsonStr []byte
					productJsonStr = []byte(`{"product_name": "Tomato"}`)
					wr = httptest.NewRecorder()
					req = httptest.NewRequest(http.MethodPost, "/api/cart/add", bytes.NewBuffer(productJsonStr))
					req.AddCookie(cookie)
					server.Handler().ServeHTTP(wr, req)
					Expect(wr.Code).To(Equal(200))

					//add to cart again
					wr = httptest.NewRecorder()
					req = httptest.NewRequest(http.MethodPost, "/api/cart/add", bytes.NewBuffer(productJsonStr))
					req.AddCookie(cookie)
					server.Handler().ServeHTTP(wr, req)
					Expect(wr.Code).To(Equal(200))

					//get cart items
					wr = httptest.NewRecorder()
					req = httptest.NewRequest(http.MethodGet, "/api/carts", nil)
					req.AddCookie(cookie)
					server.Handler().ServeHTTP(wr, req)
					Expect(wr.Code).To(Equal(200))

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
					var jsonStr = []byte(`{"username": "aditira", "password": "1234"}`)
					wr := httptest.NewRecorder()
					req := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(jsonStr))
					server.Handler().ServeHTTP(wr, req)

					Expect(wr.Code).To(Equal(200))

					var cookie *http.Cookie
					for _, c := range wr.Result().Cookies() {
						if c.Name == "token" {
							cookie = c
						}
					}

					//add to cart
					var productJsonStr []byte
					productJsonStr = []byte(`{"product_name": "Tomato"}`)

					wr = httptest.NewRecorder()
					req = httptest.NewRequest(http.MethodPost, "/api/cart/add", bytes.NewBuffer(productJsonStr))
					req.AddCookie(cookie)
					server.Handler().ServeHTTP(wr, req)
					Expect(wr.Code).To(Equal(200))

					//add to cart again
					productJsonStr = []byte(`{"product_name": "Tea"}`)
					wr = httptest.NewRecorder()
					req = httptest.NewRequest(http.MethodPost, "/api/cart/add", bytes.NewBuffer(productJsonStr))
					req.AddCookie(cookie)
					server.Handler().ServeHTTP(wr, req)
					Expect(wr.Code).To(Equal(200))

					//get cart items
					wr = httptest.NewRecorder()
					req = httptest.NewRequest(http.MethodGet, "/api/carts", nil)
					req.AddCookie(cookie)
					server.Handler().ServeHTTP(wr, req)
					Expect(wr.Code).To(Equal(200))

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
				var jsonStr = []byte(`{"username": "aditira", "password": "1234"}`)
				wr := httptest.NewRecorder()
				req := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(jsonStr))
				server.Handler().ServeHTTP(wr, req)

				Expect(wr.Code).To(Equal(200))

				var cookie *http.Cookie
				for _, c := range wr.Result().Cookies() {
					if c.Name == "token" {
						cookie = c
					}
				}

				//add to cart
				var productJsonStr []byte
				productJsonStr = []byte(`{"product_name": "Chocolate"}`)

				wr = httptest.NewRecorder()
				req = httptest.NewRequest(http.MethodPost, "/api/cart/add", bytes.NewBuffer(productJsonStr))
				req.AddCookie(cookie)
				server.Handler().ServeHTTP(wr, req)

				Expect(wr.Code).To(Equal(http.StatusNotFound))
			})
		})
	})

	Describe("cart/clear", func() {
		It("clears the cart items", func() {
			//login
			var jsonStr = []byte(`{"username": "aditira", "password": "1234"}`)
			wr := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(jsonStr))
			server.Handler().ServeHTTP(wr, req)

			Expect(wr.Code).To(Equal(200))

			var cookie *http.Cookie
			for _, c := range wr.Result().Cookies() {
				if c.Name == "token" {
					cookie = c
				}
			}

			//add to cart
			var productJsonStr []byte
			productJsonStr = []byte(`{"product_name": "Tomato"}`)
			wr = httptest.NewRecorder()
			req = httptest.NewRequest(http.MethodPost, "/api/cart/add", bytes.NewBuffer(productJsonStr))
			req.AddCookie(cookie)
			server.Handler().ServeHTTP(wr, req)
			Expect(wr.Code).To(Equal(200))

			//clear cart
			wr = httptest.NewRecorder()
			req = httptest.NewRequest(http.MethodGet, "/api/cart/clear", nil)
			req.AddCookie(cookie)
			server.Handler().ServeHTTP(wr, req)
			Expect(wr.Code).To(Equal(200))

			//get cart items
			wr = httptest.NewRecorder()
			req = httptest.NewRequest(http.MethodGet, "/api/carts", nil)
			req.AddCookie(cookie)
			server.Handler().ServeHTTP(wr, req)
			Expect(wr.Code).To(Equal(200))

			cartListSuccessResponse := api.CartListSuccessResponse{}
			json.NewDecoder(wr.Body).Decode(&cartListSuccessResponse)
			Expect(cartListSuccessResponse.CartItems).To(HaveLen(0))
		})
	})

	Describe("products", func() {
		It("returns all products", func() {
			//login
			var jsonStr = []byte(`{"username": "aditira", "password": "1234"}`)
			wr := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(jsonStr))
			server.Handler().ServeHTTP(wr, req)

			Expect(wr.Code).To(Equal(200))

			var cookie *http.Cookie
			for _, c := range wr.Result().Cookies() {
				if c.Name == "token" {
					cookie = c
				}
			}

			//get products
			wr = httptest.NewRecorder()
			req = httptest.NewRequest(http.MethodGet, "/api/products", nil)
			req.AddCookie(cookie)
			server.Handler().ServeHTTP(wr, req)
			Expect(wr.Code).To(Equal(200))

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
				var jsonStr = []byte(`{"username": "aditira", "password": "1234"}`)
				wr := httptest.NewRecorder()
				req := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(jsonStr))
				server.Handler().ServeHTTP(wr, req)

				Expect(wr.Code).To(Equal(200))

				var cookie *http.Cookie
				for _, c := range wr.Result().Cookies() {
					if c.Name == "token" {
						cookie = c
					}
				}

				//dashboard
				wr = httptest.NewRecorder()
				req = httptest.NewRequest(http.MethodGet, "/api/dashboard", nil)
				req.AddCookie(cookie)
				server.Handler().ServeHTTP(wr, req)
				Expect(wr.Code).To(Equal(200))

				dashboardSuccessResponse := api.DashboardSuccessResponse{}
				json.NewDecoder(wr.Body).Decode(&dashboardSuccessResponse)

				Expect(dashboardSuccessResponse.Username).To(Equal("aditira"))
			})

			It("returns the cart items", func() {
				//login
				var jsonStr = []byte(`{"username": "aditira", "password": "1234"}`)
				wr := httptest.NewRecorder()
				req := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(jsonStr))
				server.Handler().ServeHTTP(wr, req)

				Expect(wr.Code).To(Equal(200))

				var cookie *http.Cookie
				for _, c := range wr.Result().Cookies() {
					if c.Name == "token" {
						cookie = c
					}
				}

				//add to cart
				var productJsonStr = []byte(`{"product_name": "Tomato"}`)
				wr = httptest.NewRecorder()
				req = httptest.NewRequest(http.MethodPost, "/api/cart/add", bytes.NewBuffer(productJsonStr))
				req.AddCookie(cookie)
				server.Handler().ServeHTTP(wr, req)
				Expect(wr.Code).To(Equal(200))

				//dashboard
				wr = httptest.NewRecorder()
				req = httptest.NewRequest(http.MethodGet, "/api/dashboard", nil)
				req.AddCookie(cookie)
				server.Handler().ServeHTTP(wr, req)
				Expect(wr.Code).To(Equal(200))

				dashboardSuccessResponse := api.DashboardSuccessResponse{}
				json.NewDecoder(wr.Body).Decode(&dashboardSuccessResponse)

				Expect(dashboardSuccessResponse.CartItems).To(HaveLen(1))
				Expect(dashboardSuccessResponse.CartItems[0].ProductName).To(Equal("Tomato"))
				Expect(dashboardSuccessResponse.CartItems[0].Price).To(Equal(2200))
				Expect(dashboardSuccessResponse.CartItems[0].Quantity).To(Equal(1))
			})

			It("returns the total price", func() {
				//login
				var jsonStr = []byte(`{"username": "aditira", "password": "1234"}`)
				wr := httptest.NewRecorder()
				req := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(jsonStr))
				server.Handler().ServeHTTP(wr, req)

				Expect(wr.Code).To(Equal(200))

				var cookie *http.Cookie
				for _, c := range wr.Result().Cookies() {
					if c.Name == "token" {
						cookie = c
					}
				}

				//add to cart
				var productJsonStr []byte
				productJsonStr = []byte(`{"product_name": "Tomato"}`)
				wr = httptest.NewRecorder()
				req = httptest.NewRequest(http.MethodPost, "/api/cart/add", bytes.NewBuffer(productJsonStr))
				req.AddCookie(cookie)
				server.Handler().ServeHTTP(wr, req)
				Expect(wr.Code).To(Equal(200))

				//add to cart again
				productJsonStr = []byte(`{"product_name": "Tea"}`)
				wr = httptest.NewRecorder()
				req = httptest.NewRequest(http.MethodPost, "/api/cart/add", bytes.NewBuffer(productJsonStr))
				req.AddCookie(cookie)
				server.Handler().ServeHTTP(wr, req)
				Expect(wr.Code).To(Equal(200))

				//dashboard
				wr = httptest.NewRecorder()
				req = httptest.NewRequest(http.MethodGet, "/api/dashboard", nil)
				req.AddCookie(cookie)
				server.Handler().ServeHTTP(wr, req)
				Expect(wr.Code).To(Equal(200))

				dashboardSuccessResponse := api.DashboardSuccessResponse{}
				json.NewDecoder(wr.Body).Decode(&dashboardSuccessResponse)

				Expect(dashboardSuccessResponse.TotalPrice).To(Equal(4900))
			})

			It("calculate the change", func() {
				//login
				var jsonStr = []byte(`{"username": "aditira", "password": "1234"}`)
				wr := httptest.NewRecorder()
				req := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(jsonStr))
				server.Handler().ServeHTTP(wr, req)

				Expect(wr.Code).To(Equal(200))

				var cookie *http.Cookie
				for _, c := range wr.Result().Cookies() {
					if c.Name == "token" {
						cookie = c
					}
				}

				//add to cart
				var productJsonStr []byte
				productJsonStr = []byte(`{"product_name": "Tomato"}`)
				wr = httptest.NewRecorder()
				req = httptest.NewRequest(http.MethodPost, "/api/cart/add", bytes.NewBuffer(productJsonStr))
				req.AddCookie(cookie)
				server.Handler().ServeHTTP(wr, req)
				Expect(wr.Code).To(Equal(200))

				//dashboard
				wr = httptest.NewRecorder()
				req = httptest.NewRequest(http.MethodGet, "/api/dashboard?cash=4900", nil)
				req.AddCookie(cookie)
				server.Handler().ServeHTTP(wr, req)
				Expect(wr.Code).To(Equal(200))

				dashboardSuccessResponse := api.DashboardSuccessResponse{}
				json.NewDecoder(wr.Body).Decode(&dashboardSuccessResponse)

				Expect(dashboardSuccessResponse.TotalPrice).To(Equal(2200))
				Expect(dashboardSuccessResponse.MoneyChanges).To(Equal(2700))
			})
		})
	})

	Describe("admin", func() {
		When("username is not logged in", func() {
			It("return unauthorized", func() {
				wr := httptest.NewRecorder()
				req := httptest.NewRequest(http.MethodGet, "/api/admin/sales", nil)
				server.Handler().ServeHTTP(wr, req)

				Expect(wr.Code).To(Equal(http.StatusUnauthorized))
			})
		})
		When("user with non admin role is logged in", func() {
			It("returns forbidden", func() {
				//login
				var jsonStr = []byte(`{"username": "dito", "password": "2552"}`)
				wr := httptest.NewRecorder()
				req := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(jsonStr))
				server.Handler().ServeHTTP(wr, req)

				Expect(wr.Code).To(Equal(200))

				var cookie *http.Cookie
				for _, c := range wr.Result().Cookies() {
					if c.Name == "token" {
						cookie = c
					}
				}

				wr = httptest.NewRecorder()
				req = httptest.NewRequest(http.MethodGet, "/api/admin/sales", nil)
				req.AddCookie(cookie)
				server.Handler().ServeHTTP(wr, req)

				Expect(wr.Code).To(Equal(http.StatusForbidden))
			})
		})
		When("user with admin role is logged in", func() {
			When("request without param", func() {
				It("return all product in all time sales", func() {
					//login
					var jsonStr = []byte(`{"username": "aditira", "password": "1234"}`)
					wr := httptest.NewRecorder()
					req := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(jsonStr))
					server.Handler().ServeHTTP(wr, req)

					Expect(wr.Code).To(Equal(200))

					var cookie *http.Cookie
					for _, c := range wr.Result().Cookies() {
						if c.Name == "token" {
							cookie = c
						}
					}

					wr = httptest.NewRecorder()
					req = httptest.NewRequest(http.MethodGet, "/api/admin/sales", nil)
					req.AddCookie(cookie)
					server.Handler().ServeHTTP(wr, req)

					Expect(wr.Code).To(Equal(http.StatusOK))

					salesResponse := api.AdminResponse{}
					json.NewDecoder(wr.Body).Decode(&salesResponse)
					Expect(len(salesResponse.Sales)).To(Equal(4))
				})
			})
			When("request only product name", func() {
				It("return only requested product name", func() {
					//login
					var jsonStr = []byte(`{"username": "aditira", "password": "1234"}`)
					wr := httptest.NewRecorder()
					req := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(jsonStr))
					server.Handler().ServeHTTP(wr, req)

					Expect(wr.Code).To(Equal(200))

					var cookie *http.Cookie
					for _, c := range wr.Result().Cookies() {
						if c.Name == "token" {
							cookie = c
						}
					}

					wr = httptest.NewRecorder()
					req = httptest.NewRequest(http.MethodGet, "/api/admin/sales?product_name=Orange", nil)
					req.AddCookie(cookie)
					server.Handler().ServeHTTP(wr, req)

					Expect(wr.Code).To(Equal(http.StatusOK))

					salesResponse := api.AdminResponse{}
					json.NewDecoder(wr.Body).Decode(&salesResponse)
					Expect(len(salesResponse.Sales)).To(Equal(2))
				})
			})
			When("request only start_period", func() {
				It("return only bad request", func() {
					//login
					var jsonStr = []byte(`{"username": "aditira", "password": "1234"}`)
					wr := httptest.NewRecorder()
					req := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(jsonStr))
					server.Handler().ServeHTTP(wr, req)

					Expect(wr.Code).To(Equal(200))

					var cookie *http.Cookie
					for _, c := range wr.Result().Cookies() {
						if c.Name == "token" {
							cookie = c
						}
					}

					wr = httptest.NewRecorder()
					req = httptest.NewRequest(http.MethodGet, "/api/admin/sales?start_period=2022-05-01", nil)
					req.AddCookie(cookie)
					server.Handler().ServeHTTP(wr, req)

					Expect(wr.Code).To(Equal(http.StatusBadRequest))
				})
			})
			When("request only end_period", func() {
				It("return only bad request", func() {
					//login
					var jsonStr = []byte(`{"username": "aditira", "password": "1234"}`)
					wr := httptest.NewRecorder()
					req := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(jsonStr))
					server.Handler().ServeHTTP(wr, req)

					Expect(wr.Code).To(Equal(200))

					var cookie *http.Cookie
					for _, c := range wr.Result().Cookies() {
						if c.Name == "token" {
							cookie = c
						}
					}

					wr = httptest.NewRecorder()
					req = httptest.NewRequest(http.MethodGet, "/api/admin/sales?end_period=2022-05-01", nil)
					req.AddCookie(cookie)
					server.Handler().ServeHTTP(wr, req)

					Expect(wr.Code).To(Equal(http.StatusBadRequest))
				})
			})
			When("request start_period and end_period", func() {
				It("return returns all product within time range", func() {
					//login
					var jsonStr = []byte(`{"username": "aditira", "password": "1234"}`)
					wr := httptest.NewRecorder()
					req := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(jsonStr))
					server.Handler().ServeHTTP(wr, req)

					Expect(wr.Code).To(Equal(200))

					var cookie *http.Cookie
					for _, c := range wr.Result().Cookies() {
						if c.Name == "token" {
							cookie = c
						}
					}

					wr = httptest.NewRecorder()
					req = httptest.NewRequest(http.MethodGet, "/api/admin/sales?start_period=2022-04-30&end_period=2022-05-01", nil)
					req.AddCookie(cookie)
					server.Handler().ServeHTTP(wr, req)

					Expect(wr.Code).To(Equal(http.StatusOK))

					salesResponse := api.AdminResponse{}
					json.NewDecoder(wr.Body).Decode(&salesResponse)
					Expect(len(salesResponse.Sales)).To(Equal(3))
				})
			})
			When("request start_period, end_period, and product_name", func() {
				It("return returns requested product within time range", func() {
					//login
					var jsonStr = []byte(`{"username": "aditira", "password": "1234"}`)
					wr := httptest.NewRecorder()
					req := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBuffer(jsonStr))
					server.Handler().ServeHTTP(wr, req)

					Expect(wr.Code).To(Equal(200))

					var cookie *http.Cookie
					for _, c := range wr.Result().Cookies() {
						if c.Name == "token" {
							cookie = c
						}
					}

					wr = httptest.NewRecorder()
					req = httptest.NewRequest(http.MethodGet, "/api/admin/sales?start_period=2022-04-30&end_period=2022-05-01&product_name=Orange", nil)
					req.AddCookie(cookie)
					server.Handler().ServeHTTP(wr, req)

					Expect(wr.Code).To(Equal(http.StatusOK))

					salesResponse := api.AdminResponse{}
					json.NewDecoder(wr.Body).Decode(&salesResponse)
					Expect(len(salesResponse.Sales)).To(Equal(1))
				})
			})
		})
	})
})
