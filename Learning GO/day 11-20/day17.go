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

func loadProducts(filename string) ([]Product, error) {
	data, err := os.ReadFile(filename)

	if errors.Is(err, os.ErrNotExist) {
		return []Product{}, fmt.Errorf("файл %q не найден: %w", filename, err)
	}

	if err != nil {
		return nil, fmt.Errorf("read products file %q: %w", filename, err)
	}

	var products []Product

	if err = json.Unmarshal(data, &products); err != nil {
		return nil, fmt.Errorf("decode products file %q: %w", filename, err)
	}

	return products, nil
}

func saveProducts(filename string, products []Product) error {
	data, err := json.Marshal(products)

	if err != nil {
		return fmt.Errorf("Ошибка сериализации: %w", err)
	}
	err = os.WriteFile(filename, data, 0664)
	if err != nil {
		return fmt.Errorf("Ошибка записи: %w", err)
	}
	return nil

}

func addProducts(products []Product, newProduct Product) ([]Product, bool) {
	for _, product := range products {
		if product.Name == newProduct.Name {
			return products, false
		}
	}
	products = append(products, newProduct)
	return products, true
}

func main() {

	var products []Product
	products, err := loadProducts("./products_json17.json")
	if errors.Is(err, os.ErrNotExist) {
		products = []Product{}
	} else if err != nil {
		fmt.Println("Ошибка загрузки: ", err)
		return
	}

	newProduct := Product{
		Name:     "Monitor",
		Price:    300.00,
		Category: "Accessories",
	}

	products, added := addProducts(products, newProduct)

	if added {
		fmt.Println("Товар добавлен")
	} else {
		fmt.Println("Товар уже присутствует в списке")
	}

	err = saveProducts("./products_json17.json", products)

	if err != nil {
		fmt.Println("Ошибка сохранения", err)
		return
	}

	fmt.Printf("Успешно записаны %d продуктов \n", len(products))

}
