package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Product struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
}

func main() {
	data, err := os.ReadFile("product.json")

	if err != nil {
		fmt.Println("Ошибка чтения файла", err)
		return
	}

	var productList []Product

	err = json.Unmarshal(data, &productList)

	if err != nil {
		fmt.Println("Ошибка чтения JSON", err)
		return
	}

	for _, product := range productList {
		fmt.Println(product.Name, product.Price, product.Category)
	}
}
