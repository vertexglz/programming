package main

import "fmt"

type Transaction struct {
	Amount float64
}

func (t *Transaction) Increase(amount float64) {
	if t != nil {
		t.Amount += amount
	}
}

func (t *Transaction) GetAmount() float64 {
	if t != nil {
		return t.Amount
	}
	return 0
}

func createTransaction() *Transaction {
	newTrans := Transaction{Amount: 500}
	return &newTrans
}

func main() {
	transTest := &Transaction{Amount: 120.00}
	transTest.Increase(50.00)
	fmt.Println(transTest.GetAmount())

	var t *Transaction
	t.Increase(50)
	fmt.Println(t.GetAmount())

}
