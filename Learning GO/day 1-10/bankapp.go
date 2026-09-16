package main

import "fmt"

type User struct {
	ID     int
	Name   string
	Wallet Wallet
}

type Wallet struct {
	Balance float64
}
type Category struct {
	ID   int
	Type string
}

type Transaction struct {
	ID         int
	Amount     float64
	Type       string
	Name       string
	CategoryID int
	UserID     int
}

func (t *Transaction) reverse() {
	if t.Type == "expense" {
		t.Type = "income"
	} else if t.Type == "income" {
		t.Type = "expense"
	}
}

func (t Transaction) info() float64 {
	return t.Amount
}

func main() {
	pavel := User{
		ID:   1,
		Name: "Pavel",
		Wallet: Wallet{
			Balance: 6734.40,
		},
	}

	food := Category{
		ID:   1,
		Type: "Food",
	}

	transport := Category{
		ID:   2,
		Type: "Transport",
	}

	trans1 := Transaction{
		ID:         1,
		Amount:     120,
		Type:       "expense",
		Name:       "salad",
		CategoryID: food.ID,
		UserID:     pavel.ID,
	}
	trans2 := Transaction{
		ID:         2,
		Amount:     250,
		Type:       "income",
		Name:       "taxrent",
		CategoryID: transport.ID,
		UserID:     pavel.ID,
	}

	fmt.Println(pavel)
	fmt.Println(food)
	fmt.Println(transport)
	fmt.Println(trans1)
	fmt.Println(trans1.info())
	trans1.reverse()
	fmt.Println(trans1.Type)
	fmt.Println(trans2)

}
