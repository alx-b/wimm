package backend

import (
	"fmt"
	"strings"

	"github.com/alx-b/wimm/src/database"
)

type Wallet struct {
	db database.Database
}

/*
type WalletInterface interface {
	CloseConnection()
	GetPurchases() []database.PurchaseOutDB
	TotalPurchaseCost(purchases []database.PurchaseOutDB) float64
	GetLeftover(budget, totalPurchaseCost float64) float64
	CountTotalTags(purchases []database.PurchaseOutDB) map[string]int
	ConvertToSliceOfSliceString(purchases []database.PurchaseOutDB) [][]string
}
*/
func CreateWallet() Wallet {
	db := database.CreateDB("testing-test.db")
	return Wallet{db: db}
}

func (w *Wallet) CloseDatabaseConnection() {
	w.db.CloseConnection()
}

func (w *Wallet) GetPurchases() []database.PurchaseOutDB {
	purchases, err := w.db.GetPurchases()
	if err != nil {
		panic(err)
	}
	return purchases
}

// Get the sum of all purchases
func (w *Wallet) TotalPurchaseCost(purchases []database.PurchaseOutDB) float64 {
	sum := 0.0
	for _, purchase := range purchases {
		sum += purchase.Cost
	}
	return sum
}

// Get the difference between Budget and total purchase
func (w *Wallet) GetLeftover(budget, totalPurchaseCost float64) float64 {
	return budget - totalPurchaseCost
}

// Get the total of each category/tag
func (w *Wallet) CountTotalTags(purchases []database.PurchaseOutDB) map[string]int {
	categoriesCount := map[string]int{}
	for _, purchase := range purchases {
		categoriesCount[purchase.Tag]++
	}
	return categoriesCount
}

// Convert Purchases (slice Purchase) into slice of slice string,
// for using with the gui's table widget
func (w *Wallet) ConvertToSliceOfSliceString(purchases []database.PurchaseOutDB) [][]string {
	table := [][]string{
		{"Name", "Seller", "Tag", "Cost", "Date"},
	}
	for _, item := range purchases {
		purchaseStr := fmt.Sprintf(
			"%s,%s,%s,%.2f,%s",
			item.Name,
			item.Seller,
			item.Tag,
			item.Cost,
			item.Date,
		)
		purchaseSlice := strings.Split(purchaseStr, ",")
		table = append(table, purchaseSlice)
	}
	return table
}
