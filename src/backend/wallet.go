package backend

import (
	"fmt"
	"strings"
)

type Gain struct {
	Name     string
	From     string
	Category string
	Amount   float64
	Date     string
}

type Purchase struct {
	Name     string
	Seller   string
	Category string
	Cost     float64
	Date     string
}

type Wallet struct {
	Purchases []Purchase
}

// Get the sum of all purchases
func (w *Wallet) TotalPurchaseCost() float64 {
	sum := 0.0
	for _, purchase := range w.Purchases {
		sum += purchase.Cost
	}
	return sum
}

// Get the total of each category
func (w *Wallet) CountTotalCategories() map[string]int {
	categoriesCount := map[string]int{}
	for _, purchase := range w.Purchases {
		categoriesCount[purchase.Category]++
	}
	return categoriesCount
}

// Convert Purchases (slice Purchase) into slice of slice string,
// for using with the gui's table widget
func (w *Wallet) ConvertToSliceOfSliceString() [][]string {
	table := [][]string{
		{"Name", "Seller", "Category", "Cost", "Date"},
	}
	for _, item := range w.Purchases {
		purchaseStr := fmt.Sprintf(
			"%s,%s,%s,%.2f,%s",
			item.Name,
			item.Seller,
			item.Category,
			item.Cost,
			item.Date,
		)
		purchaseSlice := strings.Split(purchaseStr, ",")
		table = append(table, purchaseSlice)
	}
	return table
}
