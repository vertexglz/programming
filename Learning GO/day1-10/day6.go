package main

import "fmt"

type Transaction struct {
	ID     int
	Amount float64
	Name   string
}

func main() {
	transactions := map[int]Transaction{
		1: {ID: 1, Amount: 1500, Name: "Зарплата"},
		2: {ID: 2, Amount: -350, Name: "Продукты"},
		3: {ID: 3, Amount: -120, Name: "Транспорт"},
	}

	value, ok := transactions[2]

	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("Транзакция не найдена")
	}
	trans := transactions[2]
	trans.Amount = -750
	transactions[2] = trans
	fmt.Println(transactions[2])
	var balance float64

	for _, value := range transactions {
		balance += value.Amount
	}
	fmt.Println(balance)
}
