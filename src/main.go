package main

import ()

type Gain struct {
	Name     string
	From     string
	Category string
	Amount   float64
	Date     string
}

type Purchase struct {
	Name     string
	From     string
	Category string
	Cost     float64
	Date     string
}

type Wallet struct {
	Purchases []Purchase
}

func (w *Wallet) TotalPurchaseCost() float64 {
	sum := 0.0
	for _, purchase := range w.Purchases {
		sum += purchase.Cost
	}
	return sum
}
