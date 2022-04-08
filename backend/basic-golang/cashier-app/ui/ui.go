package ui

import (
	"github.com/rivo/tview"
	"github.com/ruang-guru/playground/backend/basic-golang/cashier-app/repository"
)

type UI struct {
	usersRepo       repository.UserRepository
	productsRepo    repository.ProductRepository
	cartItemRepo    repository.CartItemRepository
	transactionRepo repository.TransactionRepository
	app             *tview.Application
	pages           *tview.Pages
}

func NewUI(usersRepo repository.UserRepository, productsRepo repository.ProductRepository, cartItemRepo repository.CartItemRepository, transactionRepo repository.TransactionRepository) UI {
	return UI{
		usersRepo, productsRepo, cartItemRepo, transactionRepo, nil, nil,
	}
}

func (ui *UI) Start() {
	app := tview.NewApplication()
	ui.app = app

	pages := tview.NewPages()
	pages.AddPage("login", ui.LoginPage(), true, true)

	ui.pages = pages
	if err := app.SetRoot(pages, true).SetFocus(pages).Run(); err != nil {
		panic(err)
	}
}
