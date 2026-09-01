package main

import "fmt"

func applyDiscount(price *float64, procent float64) bool { // обязательно работать с указателем на переменную
	newPrice := *price * (1 - procent/100)
	if newPrice < 0 {
		*price = 0
		return false
	}
	*price = newPrice
	return true
}

func main() {
	var price1, procent1 float64
	price1 = 1000
	procent1 = 80
	if applyDiscount(&price1, procent1) {
		fmt.Printf("%.2f\n", price1)
	}
}
