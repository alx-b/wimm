package gui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type GraphLayoutData struct {
	Gui *GUI
}

func (d *GraphLayoutData) create1() *fyne.Container {
	monthlyPurchases := d.Gui.Wallet.GetCurrentMonthPurchases(d.Gui.Wallet.YearlyData)
	totalTags := d.Gui.Wallet.CountTotalTags(monthlyPurchases)
	graph := container.NewGridWithColumns(3)

	for x, y := range totalTags {
		bar := widget.NewProgressBar()
		bar.Min = 0.0
		bar.Max = float64(len(monthlyPurchases))
		bar.Value = float64(y)
		label := widget.NewLabel(x)
		label2 := widget.NewLabel(fmt.Sprintf("%d", y))
		graph.AddObject(label)
		graph.AddObject(label2)
		graph.AddObject(bar)
	}

	return container.NewPadded(graph)
}

func (d *GraphLayoutData) create2() *fyne.Container {
	monthlyPurchases := d.Gui.Wallet.GetCurrentMonthPurchases(d.Gui.Wallet.YearlyData)
	totalSpendingsPerTag := d.Gui.Wallet.CountTotalSpendingPerTag(monthlyPurchases)
	graph2 := container.NewGridWithColumns(3)

	for x, y := range totalSpendingsPerTag {
		bar := widget.NewProgressBar()
		bar.Min = 0.0
		bar.Max = d.Gui.Wallet.TotalPurchaseCost(monthlyPurchases)
		bar.Value = y
		label := widget.NewLabel(x)
		label2 := widget.NewLabel(fmt.Sprintf("%.2f", y))
		graph2.AddObject(label)
		graph2.AddObject(label2)
		graph2.AddObject(bar)
	}

	return container.NewPadded(graph2)
}

func CreateGraphLayout(gui *GUI) *fyne.Container {
	graphLayoutData := &GraphLayoutData{
		Gui: gui,
	}

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.NavigateBackIcon(), func() {
			graphLayoutData.Gui.CurrentWindow.SetContent(CreateMainLayout(graphLayoutData.Gui))
		}),
	)

	graph := graphLayoutData.create1()
	graph2 := graphLayoutData.create2()

	return container.NewVBox(
		toolbar,
		widget.NewSeparator(),
		widget.NewLabelWithStyle("Most 'used' tag", fyne.TextAlignLeading, fyne.TextStyle{Bold: true, TabWidth: 4}),
		graph,
		widget.NewSeparator(),
		widget.NewLabelWithStyle("Money spent per tag", fyne.TextAlignLeading, fyne.TextStyle{Bold: true, TabWidth: 4}),
		graph2,
		widget.NewSeparator(),
	)
}
