package backend

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestAddingPurchaseCost(t *testing.T) {
	t.Run("Get total purchase cost", func(t *testing.T) {
		wallet := Wallet{
			Purchases: []Purchase{
				{Name: "something1", Cost: 50.25},
				{Name: "something2", Cost: 50.25},
			},
		}
		got := wallet.TotalPurchaseCost()
		want := 100.50
		assertFloatEqual(t, got, want)

	})
	t.Run("Get total purchase cost when there is no purchase", func(t *testing.T) {
		wallet := Wallet{}
		got := wallet.TotalPurchaseCost()
		want := 0.0
		assertFloatEqual(t, got, want)
	})
}

func TestCountTotalCategories(t *testing.T) {
	t.Run("Get total of all categories", func(t *testing.T) {
		wallet := Wallet{
			Purchases: []Purchase{
				{Name: "something1", Category: "clothes"},
				{Name: "something2", Category: "food"},
				{Name: "something3", Category: "rent"},
				{Name: "something4", Category: "clothes"},
				{Name: "something5", Category: "game"},
				{Name: "something6", Category: "game"},
			},
		}
		got := wallet.CountTotalCategories()
		want := map[string]int{
			"clothes": 2,
			"food":    1,
			"game":    2,
			"rent":    1,
		}
		assertMapEqual(t, got, want)
	})
	t.Run("Get total of all categories when empty", func(t *testing.T) {
		wallet := Wallet{}
		got := wallet.CountTotalCategories()
		want := map[string]int{}
		assertMapEqual(t, got, want)
	})
}

func TestConvertToSliceOfSliceString(t *testing.T) {
	t.Run("Convert slice of struct Purchase into slice of string", func(t *testing.T) {
		wallet := Wallet{
			Purchases: []Purchase{
				{Name: "name1", Seller: "some1", Category: "clothing", Cost: 100.00, Date: "10.11.2021"},
				{Name: "name2", Seller: "some2", Category: "food", Cost: 200.00, Date: "20.11.2021"},
				{Name: "name3", Seller: "some3", Category: "rent", Cost: 300.00, Date: "25.11.2021"},
				{Name: "name4", Seller: "some4", Category: "clothing", Cost: 400.00, Date: "25.11.2021"},
			},
		}
		got := wallet.ConvertToSliceOfSliceString()
		want := [][]string{
			{"Name", "Seller", "Category", "Cost", "Date"},
			{"name1", "some1", "clothing", "100.00", "10.11.2021"},
			{"name2", "some2", "food", "200.00", "20.11.2021"},
			{"name3", "some3", "rent", "300.00", "25.11.2021"},
			{"name4", "some4", "clothing", "400.00", "25.11.2021"},
		}
		assertSliceOfSliceStringEqual(t, got, want)

	})
}

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
