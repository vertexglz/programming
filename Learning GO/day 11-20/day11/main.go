package main

import (
	"day11/wallet"
	"fmt"
)

func main() {
	myWallet, err := wallet.New(1000)
	myWallet2, err2 := wallet.New(300)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(myWallet)
	}
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(myWallet2)
	}

	err3 := wallet.Transfer(&myWallet, &myWallet2, 300)
	if err3 != nil {
		fmt.Println(err3)
	} else {
		fmt.Println(myWallet.Balance())
		fmt.Println(myWallet2.Balance())
	}

}
