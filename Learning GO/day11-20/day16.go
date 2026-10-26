package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Product struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Category string  `json:"category,omitempty"`
}

func main() {
	data, err := os.ReadFile("products_day16.json")

	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("Файл пока не создан")

		data = []byte("[]")
		err = os.WriteFile("products_day16.json", data, 0664)
		if err != nil {
			fmt.Println("Ошибка создания файла:", err)
			return
		}

		fmt.Println("Создан пустой файл")
	} else if err != nil {
		fmt.Println("Ошибка чтения:", err)
		return
	}

	var products []Product

	err = json.Unmarshal(data, &products)
	if err != nil {
		fmt.Println("Ошибка разбора JSON:", err)
		return
	}

	fmt.Printf("Загружено товаров: %d\n", len(products))
}
