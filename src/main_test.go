package main

import (
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

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
	t.Run("Get total purchase cost when there is no purchase", func(t *testing.T) {
		wallet := Wallet{}
		got := wallet.TotalPurchaseCost()
		want := 0.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}
