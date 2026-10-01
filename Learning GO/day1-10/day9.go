package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func add(a int, b int) int {
	return a + b
}

func isEnough(balance, amount float64) bool {
	if amount <= balance {
		return true
	}
	return false
}

func spend(balance, amount float64) (float64, error) {
	var err error
	if amount <= 0 {
		err = errors.New("invalid amount")
		return balance, err
	} else if amount > balance {
		err = errors.New("not enough money")
		return balance, err
	}
	return (balance - amount), err

}

func processExpenses(balance float64, amounts []float64) (float64, error) {
	var err error

	for _, value := range amounts {
		balance, err = spend(balance, value)

		if err != nil {
			return balance, err
		}
	}

	return balance, nil
}
func normalizeCategory(category string) string {
	return strings.ToLower(strings.TrimSpace(category))
}

func parseTransaction(transaction string) (string, int, error) {
	arr := strings.Split(transaction, ":")
	category := normalizeCategory(arr[0])
	if len(arr) < 2 {
		err := errors.New("Не задана цена")
		return category, 0, err
	}
	amount, err := strconv.Atoi(arr[1])
	if err != nil {
		return category, 0, errors.New("Ошибка транзакции")
	}

	return category, amount, nil
}

func parseTransactions(transaction []string) (map[string]int, error) {
	totals := make(map[string]int)
	for _, value := range transaction {
		category, amount, err := parseTransaction(value)
		if err != nil {
			return totals, err
		}
		totals[category] += amount

	}
	return totals, nil
}

func main() {
	transactions := []string{
		"Food:500",
		"Transport:200",
		"Food:300",
		"Entertainment:1000",
		"Transport:150"}

	totals, err := parseTransactions(transactions)

	if err != nil {
		fmt.Println(err)
		return
	}

	for i, value := range totals {
		fmt.Println(i, value)
	}

}
