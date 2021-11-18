package backend

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
	"time"

	"github.com/alx-b/wimm/src/database"
)

func TestAddPurchaseToDatabase(t *testing.T) {
	t.Run("Add a purchase to database", func(t *testing.T) {
		wallet := CreateWallet(":memory:")
		purchase := []string{"name1", "some1", "tag1", "100.00", "date"}
		got := wallet.AddPurchaseToDatabase(purchase)

		switch got {
		case ErrCouldNotConvertToFloat:
			t.Errorf("cannot convert to float")
		case nil:
		default:
			t.Errorf("default error")
		}
	})
}

func TestAddingPurchaseCost(t *testing.T) {
	t.Run("Get total purchase cost", func(t *testing.T) {
		wallet := Wallet{}
		purchases := []database.PurchaseOutDB{
			{Name: "something1", Cost: 50.25},
			{Name: "something2", Cost: 50.25},
		}
		got := wallet.TotalPurchaseCost(purchases)
		want := 100.50
		assertFloatEqual(t, got, want)

	})
	t.Run("Get total purchase cost when there is no purchase", func(t *testing.T) {
		wallet := Wallet{}
		purchases := []database.PurchaseOutDB{}
		got := wallet.TotalPurchaseCost(purchases)
		want := 0.0
		assertFloatEqual(t, got, want)
	})
}

func TestCountTotalTags(t *testing.T) {
	t.Run("Get total of all tags", func(t *testing.T) {
		wallet := Wallet{}
		purchases := []database.PurchaseOutDB{
			{Name: "something1", Tag: "clothes"},
			{Name: "something2", Tag: "food"},
			{Name: "something3", Tag: "rent"},
			{Name: "something4", Tag: "clothes"},
			{Name: "something5", Tag: "game"},
			{Name: "something6", Tag: "game"},
		}
		got := wallet.CountTotalTags(purchases)
		want := map[string]int{
			"clothes": 2,
			"food":    1,
			"game":    2,
			"rent":    1,
		}
		assertMapEqual(t, got, want)
	})
	t.Run("Get total of all tags when empty", func(t *testing.T) {
		wallet := Wallet{}
		purchases := []database.PurchaseOutDB{}
		got := wallet.CountTotalTags(purchases)
		want := map[string]int{}
		assertMapEqual(t, got, want)
	})
}

func TestCountSpendingsPerTag(t *testing.T) {
	t.Run("Get total spendings per tag", func(t *testing.T) {
		wallet := Wallet{}
		purchases := []database.PurchaseOutDB{
			{Name: "something1", Tag: "clothes", Cost: 1000.00},
			{Name: "something2", Tag: "food", Cost: 500.00},
			{Name: "something3", Tag: "rent", Cost: 200.00},
			{Name: "something4", Tag: "clothes", Cost: 300.00},
			{Name: "something5", Tag: "game", Cost: 100.00},
			{Name: "something6", Tag: "game", Cost: 200.00},
		}
		got := wallet.CountTotalSpendingPerTag(purchases)
		want := map[string]float64{
			"clothes": 1300.00,
			"food":    500.00,
			"rent":    200.00,
			"game":    300.00,
		}
		assertMapFloatEqual(t, got, want)

	})
}

func TestConvertToSliceOfSliceString(t *testing.T) {
	t.Run("Convert slice of struct Purchase into slice of string", func(t *testing.T) {
		wallet := Wallet{}
		purchases := []database.PurchaseOutDB{
			{Name: "name1", Seller: "some1", Tag: "clothing", Cost: 100.00, Date: "10.11.2021"},
			{Name: "name2", Seller: "some2", Tag: "food", Cost: 200.00, Date: "20.11.2021"},
			{Name: "name3", Seller: "some3", Tag: "rent", Cost: 300.00, Date: "25.11.2021"},
			{Name: "name4", Seller: "some4", Tag: "clothing", Cost: 400.00, Date: "25.11.2021"},
		}
		got := wallet.ConvertToSliceOfSliceString(purchases)
		want := [][]string{
			{"Name", "Seller", "Tag", "Cost", "Date"},
			{"name1", "some1", "clothing", "100.00", "10.11.2021"},
			{"name2", "some2", "food", "200.00", "20.11.2021"},
			{"name3", "some3", "rent", "300.00", "25.11.2021"},
			{"name4", "some4", "clothing", "400.00", "25.11.2021"},
		}
		assertSliceOfSliceStringEqual(t, got, want)

	})
}

func TestSubstractBudgetWithTotalPurchase(t *testing.T) {
	t.Run("Substract Budget with total purchase", func(t *testing.T) {
		wallet := Wallet{}
		got := wallet.GetLeftover(1999.45, 1000.00)
		want := 999.45
		assertFloatEqual(t, got, want)

	})
}

func TestGetCurrentMonthPurchases(t *testing.T) {
	t.Run("Get total of current month purchases", func(t *testing.T) {
		wallet := Wallet{Date: time.Month(11), Year: 2021}
		purchases := []database.PurchaseOutDB{
			{Name: "name1", Seller: "some1", Tag: "clothing", Cost: 100.00, Date: "2021.11.10"},
			{Name: "name2", Seller: "some2", Tag: "food", Cost: 200.00, Date: "2021.01.20"},
			{Name: "name3", Seller: "some3", Tag: "rent", Cost: 300.00, Date: "2021.12.25"},
			{Name: "name4", Seller: "some4", Tag: "clothing", Cost: 400.00, Date: "2021.11.05"},
		}
		got := wallet.GetCurrentMonthPurchases(purchases)
		want := []database.PurchaseOutDB{
			{Name: "name1", Seller: "some1", Tag: "clothing", Cost: 100.00, Date: "2021.11.10"},
			{Name: "name4", Seller: "some4", Tag: "clothing", Cost: 400.00, Date: "2021.11.05"},
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v want %v", got, want)
		}
	})
}

func TestGetMonth(t *testing.T) {
	t.Run("Get current month from time library", func(t *testing.T) {
		wallet := Wallet{}
		got := wallet.GetMonth()
		want := time.Now().Month()

		if got != want {
			t.Errorf("Got %v want %v", got, want)
		}
	})
}

func TestPrevMonth(t *testing.T) {
	t.Run("Get previous month", func(t *testing.T) {
		wallet := Wallet{Date: time.Month(6)}
		wallet.PrevMonth()
		got := wallet.Date
		want := time.Month(5)
		if got != want {
			t.Errorf("Got %v want %v", got, want)
		}
	})
	t.Run("Get previous month if month is 1 get 12", func(t *testing.T) {
		wallet := Wallet{Date: time.Month(1)}
		wallet.PrevMonth()
		got := wallet.Date
		want := time.Month(12)
		if got != want {
			t.Errorf("Got %v want %v", got, want)
		}
	})
}

func TestNextMonth(t *testing.T) {
	t.Run("Get next month", func(t *testing.T) {
		wallet := Wallet{Date: time.Month(6)}
		wallet.NextMonth()
		got := wallet.Date
		want := time.Month(7)
		if got != want {
			t.Errorf("Got %v want %v", got, want)
		}
	})
	t.Run("Get next month if month is 12 get 1", func(t *testing.T) {
		wallet := Wallet{Date: time.Month(12)}
		wallet.NextMonth()
		got := wallet.Date
		want := time.Month(1)
		if got != want {
			t.Errorf("Got %v want %v", got, want)
		}
	})
}

// ASSERT HELPERS

func assertSliceOfSliceStringEqual(t testing.TB, got, want [][]string) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertFloatEqual(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func assertMapEqual(t testing.TB, got, want map[string]int) {
	t.Helper()
	keyValueGot := []string{}
	keyValueWant := []string{}

	for keyGot, valueGot := range got {
		keyValueGot = append(keyValueGot, fmt.Sprintf("%s:%d", keyGot, valueGot))
	}
	for keyWant, valueWant := range want {
		keyValueWant = append(keyValueWant, fmt.Sprintf("%s:%d", keyWant, valueWant))
	}

	sort.Strings(keyValueGot)
	sort.Strings(keyValueWant)

	if !reflect.DeepEqual(keyValueGot, keyValueWant) {
		t.Errorf("got %v want %v", got, want)
	}
}
func assertMapFloatEqual(t testing.TB, got, want map[string]float64) {
	t.Helper()
	keyValueGot := []string{}
	keyValueWant := []string{}

	for keyGot, valueGot := range got {
		keyValueGot = append(keyValueGot, fmt.Sprintf("%s:%.2f", keyGot, valueGot))
	}
	for keyWant, valueWant := range want {
		keyValueWant = append(keyValueWant, fmt.Sprintf("%s:%.2f", keyWant, valueWant))
	}

	sort.Strings(keyValueGot)
	sort.Strings(keyValueWant)

	if !reflect.DeepEqual(keyValueGot, keyValueWant) {
		t.Errorf("got %v want %v", got, want)
	}
}
