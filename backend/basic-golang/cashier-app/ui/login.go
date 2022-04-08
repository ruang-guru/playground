package ui

import "github.com/rivo/tview"

// 1. Login Form with tview Form doc: https://github.com/rivo/tview/wiki/Form
func (ui *UI) LoginPage() *tview.Form {
	form := tview.NewForm().
		AddInputField("Username", "", 20, nil, nil).
		AddPasswordField("Password", "", 20, '*', nil)
	form.
		AddButton("Login", func() {
			username := form.GetFormItem(0).(*tview.InputField).GetText()
			password := form.GetFormItem(1).(*tview.InputField).GetText()

			_, err := ui.usersRepo.Login(username, password)
			if err != nil {
				errorModal := tview.NewModal().
					AddButtons([]string{"Ok"}).
					SetText(err.Error()).
					SetDoneFunc(
						func(buttonIndex int, buttonLabel string) {
							ui.pages.SwitchToPage("login")
							ui.app.SetRoot(ui.pages, true)
						})
				ui.app.SetRoot(errorModal, false).SetFocus(errorModal)
				return
			}
			ui.pages.AddPage("dashboard", ui.DashboardPage(0), true, true)
			ui.pages.SwitchToPage("dashboard")

		}).
		AddButton("cancel", func() {
			ui.app.Stop()
		})
	form.SetBorder(true).SetTitle(" Form Login -  (Press Tab to Navigate and Enter to Select) ").SetTitleAlign(tview.AlignLeft)
	return form
}
