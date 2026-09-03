package main

import "fmt"

type Wallet struct {
	balance float64
}

func Deposit(w *Wallet, amount float64) {
	w.balance += amount
}

func main() {
	wallet_1 := Wallet{100}
	Deposit(&wallet_1, 50)
	fmt.Print(wallet_1.balance)

}
