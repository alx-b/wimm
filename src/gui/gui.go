package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"

	"github.com/alx-b/wimm/src/backend"
)

type GUI struct {
	App           fyne.App
	CurrentWindow fyne.Window
	Wallet        *backend.Wallet
}

func CreateGUI(wallet *backend.Wallet) GUI {
	app := app.New()
	app.Settings().SetTheme(theme.DarkTheme())

	window := app.NewWindow("Where Is My Money")
	window.Resize(fyne.NewSize(500, 700))

	gui := GUI{
		App:           app,
		CurrentWindow: window,
		Wallet:        wallet,
	}

	gui.CurrentWindow.SetContent(CreateMainLayout(&gui))
	gui.CurrentWindow.ShowAndRun()

	return gui
}
