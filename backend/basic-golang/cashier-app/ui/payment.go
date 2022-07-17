package ui

import (
	"strconv"

	"github.com/rivo/tview"
)

func (ui *UI) PaymentPage(sumPrice int) *tview.Form {
	form := tview.NewForm().
		AddInputField("Cash:", "", 20, nil, nil)
	form.AddButton("Pay", func() {
		cash := form.GetFormItem(0).(*tview.InputField).GetText()
		cashInt, _ := strconv.Atoi(cash)
		ui.pages.AddAndSwitchToPage("dashboard", ui.DashboardPage(cashInt), true)
	})

	form.SetBorder(true).SetTitle(" Total Price: Rp." + strconv.Itoa(sumPrice) + " (Press Tab to Navigate and Enter to Select) ").SetTitleAlign(tview.AlignLeft)
	return form
}
