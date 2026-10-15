package main

import (
	"errors"
	"fmt"
)

type Product struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Category string  `json:"category,omitempty"`
}

type Catalog struct {
	products []Product
}

func (c *Catalog) Add(product Product) error {
	if product.Name == "" {
		return errors.New("Некорректное имя")
	}
	if product.Price <= 0 {
		return errors.New("Цена должна быть больше нуля")
	}
	for _, value := range c.products {
		if value.Name == product.Name {
			return errors.New("Данный товар уже добавлен")
		}
	}
	c.products = append(c.products, product)
	return nil

}

func (c *Catalog) Products() []Product {
	products := make([]Product, len(c.products))
	copy(products, c.products)
	return products
}

func main() {
	catalog := Catalog{}

	err := catalog.Add(Product{
		Name:     "Monitor",
		Price:    300.00,
		Category: "Accessories",
	})

	if err != nil {
		fmt.Println("Ошибка ", err)
	} else {
		fmt.Println(catalog.Products())
	}

	err = catalog.Add(Product{
		Name:     "Monitor",
		Price:    300.00,
		Category: "Accessories",
	})

	if err != nil {
		fmt.Println("Ошибка ", err)
	} else {
		fmt.Println(catalog.Products())
	}
}
