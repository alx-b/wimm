package main

import (
//"fmt"
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
