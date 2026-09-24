package main

import (
	"fmt"
)

type Product struct {
	Named
	Priced
}

type Service struct {
	Named
	Priced
}

type Priced struct {
	Price float64
}

type Named struct {
	Name string
}

type Pricer interface {
	GetPrice() float64
}

type Describer interface {
	Describe() string
}

type Info interface {
	Pricer
	Describer
}

func (p Named) Describe() string {
	return p.Name
}

func (p Priced) GetPrice() float64 {
	return p.Price
}

func (p Product) GetPrice() float64 {
	return p.Price
}

func calculateTotalPrice(items []Info) float64 {
	var totalPrice float64
	for _, item := range items {
		totalPrice += item.GetPrice()
	}
	return totalPrice
}

func printFullInfo(item Info) {
	fmt.Println(item.Describe(), item.GetPrice())
}

func main() {

	product := Product{
		Named:  Named{Name: "Keyboard"},
		Priced: Priced{Price: 28.94},
	}

	service := Service{
		Named:  Named{Name: "Taxi"},
		Priced: Priced{Price: 133.84},
	}

	printFullInfo(product)
	printFullInfo(service)

	items := []Info{
		product,
		service,
		product,
	}
	fmt.Println(calculateTotalPrice(items))
}
