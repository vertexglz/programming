package main

import (
	"day12/account"
	"fmt"
)

func main() {
	account1, err := account.New("Pavel", 1000.00)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(account1)
	}

	account2, err := account.New("Alex", 3000.00)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(account2)
	}
	err = account.Transfer(&account1, &account2, 250.00)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(account1.Balance())
		fmt.Println(account2.Balance())
	}

	fmt.Println("Hello World!")

}
