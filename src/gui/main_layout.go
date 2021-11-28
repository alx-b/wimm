package gui

import (
	"reflect"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type MainLayoutData struct {
	Gui           *GUI
	TableData     [][]string
	Table         *widget.Table
	TotalCost     binding.Float
	Leftover      binding.Float
	MonthlyBudget binding.Float
	MonthStr      binding.String
}

func CreateMainLayout(gui *GUI) *fyne.Container {
	mainLayoutData := &MainLayoutData{
		Gui:           gui,
		TotalCost:     binding.NewFloat(),
		Leftover:      binding.NewFloat(),
		MonthlyBudget: binding.NewFloat(),
		MonthStr:      binding.NewString(),
	}
	toolbar := mainLayoutData.createToolbar()
	info := mainLayoutData.createInfo()
	table := mainLayoutData.createTable()
	monthBar := mainLayoutData.createMonthBar()

	return container.NewBorder(
		container.NewVBox(
			toolbar,
			widget.NewSeparator(),
			monthBar,
			widget.NewSeparator(),
		),
		container.NewVBox(
			widget.NewSeparator(),
			info,
		),
		nil,
		nil,
		table,
	)
}

func (m *MainLayoutData) ToggleDarkAndLightTheme() {
	if reflect.DeepEqual(m.Gui.App.Settings().Theme(), theme.DarkTheme()) {
		m.Gui.App.Settings().SetTheme(theme.LightTheme())
	} else {
		m.Gui.App.Settings().SetTheme(theme.DarkTheme())
	}
}

func (m *MainLayoutData) createToolbar() *widget.Toolbar {
	return widget.NewToolbar(
		widget.NewToolbarAction(theme.ColorChromaticIcon(), func() {
			m.ToggleDarkAndLightTheme()
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			m.Gui.CurrentWindow.SetContent(m.Gui.createAddPurchaseLayout())
		}),
		widget.NewToolbarAction(theme.FileImageIcon(), func() {
			m.Gui.CurrentWindow.SetContent(CreateGraphLayout(m.Gui))
		}),
	)
}

func (m *MainLayoutData) createMonthBar() *fyne.Container {
	m.MonthStr.Set(m.Gui.Wallet.Month.String())
	displayDate := widget.NewLabelWithData(m.MonthStr)

	prevButton := widget.NewButton("<", func() {
		m.Gui.Wallet.PrevMonth()
		m.updateAll()
	})

	nextButton := widget.NewButton(">", func() {
		m.Gui.Wallet.NextMonth()
		m.updateAll()
	})

	return container.NewHBox(
		prevButton,
		displayDate,
		nextButton,
	)
}

func (m *MainLayoutData) updateAll() {
	m.MonthStr.Set(m.Gui.Wallet.Month.String())
	m.updateTable()
	m.updateInfo()
}

func (m *MainLayoutData) updateTable() {
	m.TableData = m.Gui.Wallet.ConvertToSliceOfSliceString(m.Gui.Wallet.GetCurrentMonthPurchases(m.Gui.Wallet.YearlyData))
	m.Table.Refresh()
}

func (m *MainLayoutData) updateInfo() {
	totalCost := m.Gui.Wallet.TotalPurchaseCost(m.Gui.Wallet.GetCurrentMonthPurchases(m.Gui.Wallet.YearlyData))
	m.TotalCost.Set(totalCost)
	m.Leftover.Set(m.Gui.Wallet.GetLeftover(m.Gui.Wallet.MonthlyBudget, totalCost))
}

func (m *MainLayoutData) createTable() *widget.Table {
	m.TableData = m.Gui.Wallet.ConvertToSliceOfSliceString(m.Gui.Wallet.GetCurrentMonthPurchases(m.Gui.Wallet.YearlyData))
	table := widget.NewTable(
		func() (int, int) {
			return len(m.TableData), len(m.TableData[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(m.TableData[i.Row][i.Col])
		},
	)
	for x := range m.TableData[0] {
		table.SetColumnWidth(x, 250.00)
	}
	m.Table = table
	return table
}

func (m *MainLayoutData) createInfo() *fyne.Container {
	budgetEntry := widget.NewEntry()
	m.updateInfo()
	strMonthlyBudget := binding.FloatToStringWithFormat(m.MonthlyBudget, "%.2f")
	strTotalCost := binding.FloatToStringWithFormat(m.TotalCost, "%.2f")
	strLeftover := binding.FloatToStringWithFormat(m.Leftover, "%.2f")
	return container.NewVBox(
		container.NewHBox(widget.NewLabel("Budget:"), widget.NewLabelWithData(strMonthlyBudget),
			widget.NewButton("change", func() {
				dialog.ShowForm("myForm", "confirm", "dismiss", []*widget.FormItem{widget.NewFormItem("Budget", budgetEntry)}, func(b bool) {
					if !b {
						return
					}
					floatMonthlyBudget, _ := strconv.ParseFloat(budgetEntry.Text, 64)
					m.Gui.Wallet.MonthlyBudget = floatMonthlyBudget
					m.MonthlyBudget.Set(floatMonthlyBudget)
					m.updateInfo()
				}, m.Gui.CurrentWindow)
			}),
		),
		container.NewHBox(widget.NewLabel("Spending:"), widget.NewLabelWithData(strTotalCost)),
		container.NewHBox(widget.NewLabel("Leftover:"), widget.NewLabelWithData(strLeftover)),
	)
}
