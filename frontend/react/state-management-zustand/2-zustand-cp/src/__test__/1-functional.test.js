import { render, screen, fireEvent } from "@testing-library/react"
import { act } from "react-dom/test-utils"
import App from "../App"

describe("Simplied e-commerce app with zustand test case", () => {
  it("should render the components", () => {
    render(<App />)
    const app = screen.getByLabelText("app")
    expect(app).toBeInTheDocument()
  })

  it("should add a new item", () => {
    render(<App />)
    const openModalNewItem = screen.getByLabelText("add-item-modal-open")

    act(() => {
      openModalNewItem.click()
    })
    const addItemIdInput = screen.getByLabelText("add-item-id-input")
    expect(addItemIdInput).toBeInTheDocument()
    const addItemNameInput = screen.getByLabelText("add-item-name-input")
    expect(addItemNameInput).toBeInTheDocument()
    const addItemImageInput = screen.getByLabelText("add-item-image-input")
    expect(addItemImageInput).toBeInTheDocument()
    const addItemPriceInput = screen.getByLabelText("add-item-price-input")
    expect(addItemPriceInput).toBeInTheDocument()
    const addItemStockInput = screen.getByLabelText("add-item-stock-input")
    expect(addItemStockInput).toBeInTheDocument()
    const addItemButton = screen.getByLabelText("add-item-button")
    expect(addItemButton).toBeInTheDocument()

    fireEvent.change(addItemIdInput, { target: { value: "ID-12345" } })
    fireEvent.change(addItemNameInput, { target: { value: "Bread" } })
    fireEvent.change(addItemImageInput, {
      target: {
        value:
          "https://assets.bonappetit.com/photos/5c62e4a3e81bbf522a9579ce/1:1/w_2240,c_limit/milk-bread.jpg",
      },
    })
    fireEvent.change(addItemPriceInput, { target: { value: "1000" } })
    fireEvent.change(addItemStockInput, { target: { value: "10" } })

    act(() => {
      addItemButton.click()
    })

    const itemName = screen.getByLabelText("item-name")
    expect(itemName).toBeInTheDocument()
    expect(itemName).toHaveTextContent("Bread")

    const itemPrice = screen.getByLabelText("item-price")
    expect(itemPrice).toBeInTheDocument()
    expect(itemPrice).toHaveTextContent("Rp. 1000")

    const itemStock = screen.getByLabelText("item-stock")
    expect(itemStock).toBeInTheDocument()
    expect(itemStock).toHaveTextContent("10")
  })

  it("should add new multiple items", () => {
    const items = [
      {
        id: 2354435,
        name: "MilkyBread",
        image:
          "https://assets.bonappetit.com/photos/5c62e4a3e81bbf522a9579ce/1:1/w_2240,c_limit/milk-bread.jpg",
        price: 2.99,
        stock: 10,
      },
      {
        id: 2354436,
        name: "Milk",
        image:
          "https://assets.bonappetit.com/photos/5c62e4a3e81bbf522a9579ce/1:1/w_2240,c_limit/milk-bread.jpg",
        price: 2.99,
        stock: 10,
      },
    ]

    render(<App />)
    const openModalNewItem = screen.getByLabelText("add-item-modal-open")

    for (let i = 0; i < items.length; i++) {
      act(() => {
        openModalNewItem.click()
      })
      const addItemIdInput = screen.getByLabelText("add-item-id-input")
      expect(addItemIdInput).toBeInTheDocument()
      const addItemNameInput = screen.getByLabelText("add-item-name-input")
      expect(addItemNameInput).toBeInTheDocument()
      const addItemImageInput = screen.getByLabelText("add-item-image-input")
      expect(addItemImageInput).toBeInTheDocument()
      const addItemPriceInput = screen.getByLabelText("add-item-price-input")
      expect(addItemPriceInput).toBeInTheDocument()
      const addItemStockInput = screen.getByLabelText("add-item-stock-input")
      expect(addItemStockInput).toBeInTheDocument()
      const addItemButton = screen.getByLabelText("add-item-button")
      expect(addItemButton).toBeInTheDocument()

      fireEvent.change(addItemIdInput, { target: { value: items[i].id } })
      fireEvent.change(addItemNameInput, { target: { value: items[i].name } })
      fireEvent.change(addItemImageInput, { target: { value: items[i].image } })
      fireEvent.change(addItemPriceInput, { target: { value: items[i].price } })
      fireEvent.change(addItemStockInput, { target: { value: items[i].stock } })

      act(() => {
        addItemButton.click()
      })
    }

    const itemsName = screen.getAllByLabelText("item-name")
    expect(itemsName[0]).toBeInTheDocument()
    expect(itemsName).toHaveLength(3)
    expect(itemsName[2]).toHaveTextContent("Milk")
  })

  it("should add item to cart", () => {
    render(<App />)
    const addToCartButtons = screen.getAllByLabelText("item-add-to-cart-button")
    expect(addToCartButtons).toHaveLength(3)

    act(() => {
        addToCartButtons[0].click()
        }
    )

    const cartOpenButton = screen.getByLabelText("cart-open-button")
    expect(cartOpenButton).toBeInTheDocument()
    act(() => {
        cartOpenButton.click()
        }
    )

    const cartItemsName = screen.getAllByLabelText("cart-item-name")

    expect(cartItemsName).toHaveLength(1)
    expect(cartItemsName[0]).toHaveTextContent("Bread")
  })

    it("should add multiple items to cart", () => {
        render(<App />)
        const addToCartButtons = screen.getAllByLabelText("item-add-to-cart-button")
        expect(addToCartButtons).toHaveLength(3)

        act(() => {
            addToCartButtons[0].click()
            addToCartButtons[1].click()
            }
        )

        const cartOpenButton = screen.getByLabelText("cart-open-button")
        expect(cartOpenButton).toBeInTheDocument()
        act(() => {
            cartOpenButton.click()
            }
        )

        const cartItemsName = screen.getAllByLabelText("cart-item-name")

        expect(cartItemsName).toHaveLength(2)
        expect(cartItemsName[0]).toHaveTextContent("Bread")
        expect(cartItemsName[1]).toHaveTextContent("MilkyBread")
    })

    it("should remove item from cart", () => {
        render(<App />)

        const cartOpenButton = screen.getByLabelText("cart-open-button")

        act(() => {
            cartOpenButton.click()
            }
        )

        const removeCartItemButtons = screen.getAllByLabelText("cart-item-remove-button")
        expect(removeCartItemButtons).toHaveLength(2)

        act(() => {
            removeCartItemButtons[0].click()
            }
        )

        const cartItemsName = screen.getAllByLabelText("cart-item-name")
        expect(cartItemsName).toHaveLength(1)
    })

    it('should change item quantity in cart', () => {
        render(<App />)

        const cartOpenButton = screen.getByLabelText("cart-open-button")

        act(() => {
            cartOpenButton.click()
            }
        )

        const cartItemsQuantity = screen.getAllByLabelText("cart-item-quantity")
        expect(cartItemsQuantity).toHaveLength(1)

        const incrementCartItemButtons = screen.getAllByLabelText("cart-item-increment-button")
        expect(incrementCartItemButtons).toHaveLength(1)

        for (let i=0; i<5; i++) {
            act(() => {
                incrementCartItemButtons[0].click()
                }
            )
        }
        expect(cartItemsQuantity[0]).toHaveTextContent("6")

        const decrementCartItemButtons = screen.getAllByLabelText("cart-item-decrement-button")
        expect(decrementCartItemButtons).toHaveLength(1)

        for (let i=0; i<2; i++) {
            act(() => {
                decrementCartItemButtons[0].click()
                }
            )
        }
        expect(cartItemsQuantity[0]).toHaveTextContent("4")

    })

    it("should limit cart item quantity from 1 to stock", () => {
        render(<App />)

        const cartOpenButton = screen.getByLabelText("cart-open-button")

        act(() => {
            cartOpenButton.click()
            }
        )

        const cartItemsQuantity = screen.getAllByLabelText("cart-item-quantity")
        expect(cartItemsQuantity).toHaveLength(1)

        const incrementCartItemButtons = screen.getAllByLabelText("cart-item-increment-button")

        for (let i=0; i<10; i++) {
            act(() => {
                incrementCartItemButtons[0].click()
                }
            )
        }
        expect(cartItemsQuantity[0]).toHaveTextContent("10")


        const decrementCartItemButtons = screen.getAllByLabelText("cart-item-decrement-button")
        expect(decrementCartItemButtons).toHaveLength(1)

        for (let i=0; i<13; i++) {
            act(() => {
                decrementCartItemButtons[0].click()
                }
            )
        }
        expect(cartItemsQuantity[0]).toHaveTextContent("1")
    })
})
