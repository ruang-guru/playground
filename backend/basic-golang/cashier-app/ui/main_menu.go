package ui

import (
	"github.com/rivo/tview"
)

func (ui *UI) MainMenuPage(username string, totalPrice int) *tview.Modal {
	modal := tview.NewModal().
		SetText("Select Action Cashier").
		AddButtons([]string{"Order", "Pay", "Exit"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Order" {
				ui.pages.AddAndSwitchToPage("productList", ui.ProductListPage(), true)
			} else if buttonLabel == "Pay" {
				ui.pages.AddAndSwitchToPage("payment", ui.PaymentPage(totalPrice), true)
			} else {
				ui.usersRepo.Logout(username)

				exitModal := tview.NewModal().
					AddButtons([]string{"Ok"}).SetText("Thank you for using the Cashier App!").SetDoneFunc(func(buttonIndex int, buttonLabel string) {
					ui.cartItemRepo.ResetCartItems()

					ui.app.Stop()
				})
				exitModal.SetRect(0, 0, 100, 20)
				ui.app.SetRoot(exitModal, false).SetFocus(exitModal)
			}
		})
	return modal
}
