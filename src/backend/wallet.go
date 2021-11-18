package backend

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/alx-b/wimm/src/database"
)

const (
	ErrCouldNotConvertToFloat = WalletError("Field needs to be a number.")
)

type WalletError string

func (e WalletError) Error() string {
	return string(e)
}

type Wallet struct {
	Year int
	Date time.Month
	db   database.Database
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

// Create database and return it as Wallet struct
func CreateWallet(filePath string) Wallet {
	db := database.CreateDB(filePath)
	return Wallet{db: db, Date: time.Now().Month(), Year: time.Now().Year()}
}

// Close connection to database
func (w *Wallet) CloseDatabaseConnection() error {
	return w.db.CloseConnection()
}

// Convert string to float and send data to database
func (w *Wallet) AddPurchaseToDatabase(purchase []string) error {
	convCost, err := strconv.ParseFloat(purchase[3], 64)

	if err != nil {
		return ErrCouldNotConvertToFloat
	}

	convPurchase := database.PurchaseInDB{
		Name:   purchase[0],
		Seller: purchase[1],
		Tag:    purchase[2],
		Cost:   convCost,
		Date:   purchase[4],
	}
	err = w.db.AddPurchase(convPurchase)

	return err
}

// Get all purchases from Database
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

// Get total spending of each category/tag
func (w *Wallet) CountTotalSpendingPerTag(purchases []database.PurchaseOutDB) map[string]float64 {
	tagSpendings := map[string]float64{}
	for _, purchase := range purchases {
		tagSpendings[purchase.Tag] += purchase.Cost
	}
	return tagSpendings
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

// Get only purchases of a specific month out of the yearly purchase
func (w *Wallet) GetCurrentMonthPurchases(yearlyPurchases []database.PurchaseOutDB) []database.PurchaseOutDB {
	monthlyPurchases := []database.PurchaseOutDB{}
	for _, purchase := range yearlyPurchases {
		if strings.Contains(purchase.Date, fmt.Sprintf(".%d.", w.Date)) {
			monthlyPurchases = append(monthlyPurchases, purchase)
		}
	}
	return monthlyPurchases
}

// Call from the time library, time .Now() .Month()
func (w *Wallet) GetMonth() time.Month {
	return time.Now().Month()
}

// Cycle throught month backward
func (w *Wallet) PrevMonth() {
	if w.Date > 1 {
		w.Date--
	} else {
		w.Date = time.Month(12)
	}
}

// Cycle throught month forward
func (w *Wallet) NextMonth() {
	if w.Date < 12 {
		w.Date++
	} else {
		w.Date = time.Month(1)
	}
}
