package ui

import (
	"strconv"

	"github.com/rivo/tview"
)

func (ui *UI) ProductListPage() *tview.List {
	list := tview.NewList()

	products, err := ui.productsRepo.SelectAll()
	if err != nil {
		panic(err)
	}
	for _, product := range products {
		p := product
		list.AddItem(product.ProductName+": Rp."+strconv.Itoa(product.Price), product.Category, '*', func() {
			ui.cartItemRepo.Add(p)
			ui.pages.AddAndSwitchToPage("dashboard", ui.DashboardPage(0), true)
		})
	}

	list.SetBorder(true).SetTitle("Shopping List - Press ENTER to Select Product").SetTitleAlign(tview.AlignLeft)
	return list
}
