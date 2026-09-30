package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	text := " Food,Transport,food,Entertainment,transport,Food "
	arr := strings.Split(text, ",")
	count := make(map[string]int)

	for _, value := range arr {
		value = strings.ToLower(strings.TrimSpace(value))
		count[value]++
	}

	for i, value := range count {
		fmt.Println(i, value)
	}

	transactions := []string{
		"Food:500",
		"Transport:200",
		"Food:300",
		"Entertainment:1000",
		"Transport:150",
		"Food : abc",
	}

	totals := make(map[string]int)
	for _, transaction := range transactions {
		value := strings.Split(transaction, ":")
		amount, err := strconv.Atoi(value[1])
		if err != nil {
			fmt.Println("Ошибка в транзакции:", transaction)
			continue
		}
		totals[strings.ToLower(strings.TrimSpace(value[0]))] += amount
	}

	for category, value := range totals {
		fmt.Println(category, value)
	}

}
