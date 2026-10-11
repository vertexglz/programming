package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Product struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Category string  `json:"category,omitempty"`
}

func main() {
	product1 := Product{
		Name:     "Laptop",
		Price:    1200.00,
		Category: "Electronics",
	}

	product2 := Product{
		Name:  "Mouse",
		Price: 25.00,
	}

	products := []Product{product1, product2}

	data, err := json.Marshal(products)

	if err != nil {
		fmt.Println("Ошибка сериализации", err)
		return
	}

	err = os.WriteFile("product.json", data, 0644)

	if err != nil {
		fmt.Println("Ошибка записи в файл", err)
		return
	}

	fmt.Println("Файл успешно записан")

	data, err = os.ReadFile("product.json")

	if err != nil {
		fmt.Println("Ошибка чтения", err)
		return
	}

	var loadProducts []Product

	err = json.Unmarshal(data, &loadProducts)

	if err != nil {
		fmt.Println("Ошибка десериализации", err)
		return
	}

	for _, product := range loadProducts {
		fmt.Println(product.Name, product.Price, product.Category)
	}
}
