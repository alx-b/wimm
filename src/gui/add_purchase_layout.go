package gui

import (
	"fmt"
	"regexp"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var r, _ = regexp.Compile("^\\d{4}\\.\\d{2}\\.\\d{2}$")

func (g *GUI) createAddPurchaseLayout() *fyne.Container {
	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Enter purchase name")
	sellerEntry := widget.NewEntry()
	sellerEntry.SetPlaceHolder("Enter seller")
	tagSelectEntry := widget.NewSelectEntry(g.Wallet.GetAllTags())
	tagSelectEntry.SetPlaceHolder("Enter or select Tag/Category")
	costEntry := widget.NewEntry()
	costEntry.SetPlaceHolder("Enter purchase cost")
	dateEntry := widget.NewEntry()
	dateEntry.SetPlaceHolder("Enter purchase date YYYY.MM.DD")
	messageLabel := widget.NewLabel("")

	submitButton := widget.NewButton("Add", func() {
		if nameEntry.Text == "" {
			messageLabel.SetText("Purchase name field cannot be blank.")
			return
		}

		if !r.MatchString(dateEntry.Text) {
			messageLabel.SetText("Date should be written YYYY.MM.DD")
			return
		}

		monthInt, _ := strconv.Atoi(dateEntry.Text[5:7])
		if monthInt > 12 || monthInt < 1 {
			messageLabel.SetText("Months are from 01 to 12")
			return
		}

		dayInt, _ := strconv.Atoi(dateEntry.Text[8:10])
		if dayInt > 31 || dayInt < 1 {
			messageLabel.SetText("Days are from 01 to 31")
			return
		}

		err := g.Wallet.AddPurchaseToDatabase(
			[]string{
				nameEntry.Text,
				sellerEntry.Text,
				tagSelectEntry.Text,
				costEntry.Text,
				dateEntry.Text,
			},
		)

		if err != nil {
			messageLabel.SetText(fmt.Sprintf("Error: %s", err))
		}

		nameEntry.SetText("")
		sellerEntry.SetText("")
		tagSelectEntry.SetText("")
		costEntry.SetText("")
		dateEntry.SetText("")
		messageLabel.SetText("")

		g.Wallet.YearlyData = g.Wallet.GetPurchases()
	})

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.NavigateBackIcon(), func() {
			g.CurrentWindow.SetContent(CreateMainLayout(g))
		}),
	)

	form := container.NewVBox(
		nameEntry,
		sellerEntry,
		tagSelectEntry,
		costEntry,
		dateEntry,
		submitButton,
		messageLabel,
	)

	return container.NewBorder(toolbar, nil, nil, nil, form)
}
