package main

import "fmt"

type Product struct {
	Name    string
	Price   float64
	InStock bool
}

func CalculateInventory(products []Product) (totalValue float64, availableCount int) {
	for _, value := range products {
		if value.InStock {
			totalValue += value.Price
			availableCount++
		}
	}
	return totalValue, availableCount

}

func main() {
	items := []Product{
		{"Ноутбук", 23645.54, true},
		{"Мышь", 254.12, true},
		{"Клавиатура", 277.12, false},
	}

	fmt.Println(CalculateInventory(items))
}
