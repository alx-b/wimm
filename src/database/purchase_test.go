package database

import (
	//	"fmt"
	"reflect"
	"testing"
)

func TestGetPurchases(t *testing.T) {
	db := CreateDB(":memory:")
	defer db.CloseConnection()
	db.AddPurchase(PurchaseInDB{Name: "name1", Seller: "", Tag: "", Cost: 25.25, Date: ""})
	db.AddPurchase(PurchaseInDB{Name: "name2", Seller: "", Tag: "", Cost: 25.50, Date: ""})

	t.Run("Get all purchases from database", func(t *testing.T) {
		got, err := db.GetPurchases()
		want := []PurchaseOutDB{
			{
				Id:   1,
				Name: "name1",
				Cost: 25.25,
			},
			{
				Id:   2,
				Name: "name2",
				Cost: 25.50,
			},
		}
		if err != nil {
			t.Errorf("%s", err.Error())
		}
		assertPurchaseListDeepEqual(t, got, want)
	})
}

func TestAddPurchase(t *testing.T) {
	db := CreateDB(":memory:")
	defer db.CloseConnection()

	t.Run("Add Purchase to database", func(t *testing.T) {
		err1 := db.AddPurchase(PurchaseInDB{Name: "name1", Seller: "", Tag: "", Cost: 25.25, Date: ""})
		err2 := db.AddPurchase(PurchaseInDB{Name: "name2", Seller: "", Tag: "", Cost: 25.50, Date: ""})
		assertNoError(t, err1)
		assertNoError(t, err2)
		got, err := db.GetPurchases()
		want := []PurchaseOutDB{
			{
				Id:   1,
				Name: "name1",
				Cost: 25.25,
			},
			{
				Id:   2,
				Name: "name2",
				Cost: 25.50,
			},
		}
		assertNoError(t, err)
		assertPurchaseListDeepEqual(t, got, want)
	})

	t.Run("Add Purchase to database with some blank", func(t *testing.T) {
		err := db.AddPurchase(PurchaseInDB{Name: "name4", Cost: 3325.12})
		assertNoError(t, err)
		got, err := db.GetPurchases()
		want := []PurchaseOutDB{
			{
				Id:   1,
				Name: "name1",
				Cost: 25.25,
			},
			{
				Id:   2,
				Name: "name2",
				Cost: 25.50,
			},
			{
				Id:   3,
				Name: "name4",
				Cost: 3325.12,
			},
		}
		assertNoError(t, err)
		assertPurchaseListDeepEqual(t, got, want)
	})
}

func TestGetPurchaseOfSpecificYear(t *testing.T) {
	db := CreateDB(":memory:")
	defer db.CloseConnection()
	db.AddPurchase(PurchaseInDB{Name: "name1", Seller: "", Tag: "", Cost: 25.25, Date: "2020.10.10"})
	db.AddPurchase(PurchaseInDB{Name: "name2", Seller: "", Tag: "", Cost: 30.50, Date: "2019.10.11"})
	db.AddPurchase(PurchaseInDB{Name: "name3", Seller: "", Tag: "", Cost: 25.50, Date: "2020.11.11"})

	t.Run("Get Purchase of specific year", func(t *testing.T) {
		got, err := db.GetPurchasesOfSpecificYear("2020")
		want := []PurchaseOutDB{
			{
				Id:   1,
				Name: "name1",
				Cost: 25.25,
				Date: "2020.10.10",
			},
			{
				Id:   3,
				Name: "name3",
				Cost: 25.50,
				Date: "2020.11.11",
			},
		}
		assertNoError(t, err)
		assertPurchaseListDeepEqual(t, got, want)
	})
}

func assertNoError(t testing.TB, err error) {
	if err != nil {
		t.Errorf("%s", err.Error())
	}
}

func assertPurchaseListDeepEqual(t testing.TB, got, want []PurchaseOutDB) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v want %#v", got, want)
	}
}
