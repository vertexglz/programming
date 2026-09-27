package main

import "fmt"

type User struct {
	ID   int
	Name string
}

func main() {

	users := map[int]User{
		1: {ID: 1, Name: "Pavel"},
		2: {ID: 2, Name: "Alex"},
	}

	fmt.Println(users[1].Name)

	products := map[int]float64{
		101: 1500,
		102: 2300,
		103: 750,
	}

	fmt.Println(products[102])
	product, ok := products[999]
	if ok {
		fmt.Println(product)
	} else {
		fmt.Println("Товар не найден")
	}

	for id, value := range products {
		fmt.Println(id, value)
	}

}
