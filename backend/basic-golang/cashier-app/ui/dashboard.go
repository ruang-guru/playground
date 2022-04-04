package ui

import (
	"strconv"

	"github.com/rivo/tview"
)

func (ui *UI) DashboardPage(cash int) *tview.Grid {
	username, err := ui.usersRepo.FindLoggedinUser()
	if err != nil {
		panic(err)
	}

	createText := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text)
	}

	cartItems := func() tview.Primitive {
		cartItems, err := ui.cartItemRepo.SelectAll()
		if err != nil {
			panic(err)
		}
		list := tview.NewList()
		list.AddItem("Shopping cart", "List Product, Price & Quantity:", '*', nil)
		for _, product := range cartItems {
			list.AddItem(product.ProductName+": Rp."+strconv.Itoa(product.Price)+" Qty:"+strconv.Itoa(product.Quantity), product.Category, '-', nil)
		}

		return list
	}

	sumPrice, err := ui.cartItemRepo.TotalPrice()
	if err != nil {
		panic(err)
	}

	moneyChanges, err := ui.transactionRepo.Pay(cash)
	if err != nil {
		panic(err)
	}

	grid := tview.NewGrid().
		SetRows(3, 0, 3).
		SetColumns(30, 0, 30).
		SetBorders(true).
		AddItem(createText("Cashier Apps"), 0, 0, 1, 3, 0, 0, false)

	account := createText("Username: " + *username + "\nCash: Rp." + strconv.Itoa(cash) + "\nMoney Changes: Rp." + strconv.Itoa(moneyChanges))
	totalPrice := createText("Total Price:\nRp. " + strconv.Itoa(sumPrice))
	grid.AddItem(account, 1, 0, 1, 1, 0, 100, false).
		AddItem(cartItems(), 1, 1, 1, 1, 0, 100, false).
		AddItem(totalPrice, 1, 2, 1, 1, 0, 100, false)

	btnAction := tview.NewButton("Press ENTER to Action").SetSelectedFunc(func() {
		ui.pages.AddAndSwitchToPage("mainMenu", ui.MainMenuPage(*username, sumPrice), true)
	})
	grid.AddItem(btnAction, 2, 0, 1, 3, 0, 0, true)

	return grid
}
