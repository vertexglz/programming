package main

import (
	"day11/wallet"
	"fmt"
)

func main() {
	myWallet, err := wallet.New(1000)
	myWallet2, err2 := wallet.New(-1000)

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

}
