package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Product struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Category string  `json:"category,omitempty"`
}

type Catalog struct {
	products []Product
}

var (
	ErrProductNotFound      = errors.New("Товар не найден")
	ErrProductAlreadyExists = errors.New("Товар уже существует")
	ErrEmptyProductName     = errors.New("Название товара не может быть пустым")
	ErrInvalidProductPrice  = errors.New("Цена товара должна быть больше 0")
)

// validate проверяет корректность товара перед добавлением и обновлением.
func validate(p Product) error {
	if strings.TrimSpace(p.Name) == "" {
		return ErrEmptyProductName
	}
	if p.Price <= 0 {
		return ErrInvalidProductPrice
	}
	return nil
}

// indexOf возвращает индекс товара по названию или -1, если товара нет.
func (c *Catalog) indexOf(name string) int {
	for i, p := range c.products {
		if p.Name == name {
			return i
		}
	}
	return -1
}

// Add добавляет товар в каталог; возвращает ErrProductAlreadyExists при дубликате.
func (c *Catalog) Add(product Product) error {
	if err := validate(product); err != nil {
		return err
	}
	if c.indexOf(product.Name) != -1 {
		return ErrProductAlreadyExists
	}
	c.products = append(c.products, product)
	return nil
}

// Get возвращает товар по названию или ErrProductNotFound.
func (c *Catalog) Get(name string) (Product, error) {
	if i := c.indexOf(name); i != -1 {
		return c.products[i], nil
	}
	return Product{}, ErrProductNotFound
}

// Update заменяет товар с названием name на update.
// Обновление с переименованием на существующий товар запрещено.
func (c *Catalog) Update(name string, update Product) error {
	index := c.indexOf(name)
	if index == -1 {
		return ErrProductNotFound
	}
	if err := validate(update); err != nil {
		return err
	}
	if i := c.indexOf(update.Name); i != -1 && i != index {
		return ErrProductAlreadyExists
	}
	c.products[index] = update
	return nil
}

// Remove удаляет товар по названию.
func (c *Catalog) Remove(name string) error {
	index := c.indexOf(name)
	if index == -1 {
		return ErrProductNotFound
	}
	c.products = append(c.products[:index], c.products[index+1:]...)
	return nil
}

func (c *Catalog) Products() []Product {
	products := []Product{}
	for _, value := range c.products {
		products = append(products, value)
	}
	return products
}

func loadProducts() ([]Product, error) {
	data, err := os.ReadFile("products.json")
	var products []Product
	if errors.Is(err, os.ErrNotExist) {
		return products, nil
	}
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &products)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func saveProducts(products []Product) error {
	data, err := json.MarshalIndent(products, "", "   ")
	if err != nil {
		return err
	}
	return os.WriteFile("products.json", data, 0664)

}

func NewCatalog(products []Product) (*Catalog, error) {
	catalog := &Catalog{}

	for _, value := range products {
		err := catalog.Add(value)
		if err != nil {
			return nil, err
		}
	}
	return catalog, nil
}

func main() {
	products, err := loadProducts()
	if err != nil {
		fmt.Println(err)
		return
	}
	catalog, err := NewCatalog(products)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = catalog.Get("Mouse")
	if errors.Is(err, ErrProductNotFound) {
		err = catalog.Add(Product{Name: "Mouse", Price: 120.00, Category: "Accessories"})
	} else if err != nil {
		fmt.Println(err)
		return
	} else {
		err = catalog.Update("Mouse", Product{Name: "Mouse", Price: 130.00, Category: "Accessories"})
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	err = saveProducts(catalog.Products())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Операция завершилась успешно")
}
