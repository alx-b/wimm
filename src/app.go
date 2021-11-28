package app

import (
	"github.com/alx-b/wimm/src/backend"
	"github.com/alx-b/wimm/src/gui"
)

func Run() {
	wallet := backend.CreateWallet("testing-test.db")
	defer wallet.CloseDatabaseConnection()
	gui.CreateGUI(&wallet)
}
