package main

import (
	"errors"
	"fmt"
)

type Wallet struct {
	Balance float64
}

type User struct {
	Name   string
	Wallet Wallet
}

func (u *User) Spend(amount float64) error {
	if u == nil {
		return errors.New("Пользователь не найден")
	}
	if amount <= 0 {
		return errors.New("Недопустимая сумма")
	}
	if amount > u.Wallet.Balance {
		return errors.New("Недостаточно средств")
	} else {
		u.Wallet.Balance -= amount
		return nil
	}
}

func main() {
	Pavel := User{
		Name:   "Pavel",
		Wallet: Wallet{1000},
	}

	fmt.Println(Pavel.Spend(0))
	fmt.Println(Pavel.Spend(-100))
	fmt.Println(Pavel.Spend(1500))
	fmt.Println(Pavel.Spend(250))
	fmt.Println(Pavel.Wallet.Balance)

}
