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

var ErrProductNotFound = errors.New("Товар не найден")
var ErrProductAlreadyExists = errors.New("Товар уже существует")

func (c *Catalog) Add(product Product) error {
	if product.Name == "" {
		return errors.New("Название товара не может быть пустым")
	}
	if product.Price <= 0 {
		return errors.New("Цена товара должна быть больше 0")
	}
	for _, value := range c.products {
		if value.Name == product.Name {
			return ErrProductAlreadyExists
		}
	}
	c.products = append(c.products, product)
	return nil

}

func (c *Catalog) Get(name string) (Product, error) {
	for _, value := range c.products {
		if value.Name == name {
			return value, nil
		}
	}
	return Product{}, ErrProductNotFound
}

func (c *Catalog) Update(name string, update Product) error {
	index := -1
	for i, value := range c.products {
		if value.Name == name {
			index = i
			break
		}

	}
	if index == -1 {
		return ErrProductNotFound
	}
	if update.Name == "" {
		return errors.New("Имя товара не соответствует описанию")
	}
	if update.Price <= 0 {
		return errors.New("Цена товара должна быть больше 0")
	}
	for i, value := range c.products {
		if value.Name == update.Name && i != index {
			return ErrProductAlreadyExists
		}
	}
	c.products[index] = update
	return nil

}

func (c *Catalog) Remove(name string) error {
	index := -1

	for i, value := range c.products {
		if value.Name == name {
			index = i
			break
		}
	}

	if index == -1 {
		return ErrProductNotFound
	}

	c.products = append(c.products[:index], c.products[index+1:]...)
	return nil
}

func main() {
	catalog := Catalog{}

	err := catalog.Add(Product{Name: "Monitor", Price: 300, Category: "Accessories"})

	if err != nil {
		fmt.Println("Ошибка добавления:", err)
		return
	}

	product, err := catalog.Get("Monitor")

	if err != nil {
		fmt.Println("Ошибка поиска", err)
		return
	}
	fmt.Println(product)

	_, err = catalog.Get("Laptop")
	if errors.Is(err, ErrProductNotFound) {
		fmt.Println("Laptop: ", err)
	}

	err = catalog.Update("Monitor", Product{
		Name:     "Monitor",
		Price:    280,
		Category: "Accessories",
	})
	if err != nil {
		fmt.Println("Ошибка обновления", err)
		return
	}

	product, err = catalog.Get("Monitor")

	if err != nil {
		fmt.Println("Ошибка поиска", err)
		return
	}
	fmt.Println(product)

	err = catalog.Remove("Monitor")
	if err != nil {
		fmt.Println("Ошибка удаления:", err)
	} else {
		fmt.Println("Monitor удалён")
	}

	err = catalog.Remove("Laptop")
	if errors.Is(err, ErrProductNotFound) {
		fmt.Println("Laptop не найден")
	}

}
