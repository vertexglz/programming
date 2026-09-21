package main

import "fmt"

type BankAccount struct {
	Owner   string
	Balance float64
}

func (a *BankAccount) Deposit(amount float64) {
	if amount > 0 {
		a.Balance += amount
	} else {
		fmt.Println("Не возможно пополнить баланс на отрицательную сумму")
	}
}

func (a *BankAccount) Withdraw(amount float64) {
	if amount < 0 {
		fmt.Println("Не возможно снять отрицательную сумму")
	} else if amount > a.Balance {
		fmt.Println("Недостаточно средств для снятия")
	} else {
		a.Balance -= amount
	}
}

func (a BankAccount) GetBalance() float64 {
	return a.Balance
}

func (a BankAccount) PrintInfo() {
	fmt.Println(a)
}

func main() {
	acc1 := BankAccount{Owner: "Pavel",
		Balance: 6743.00}

	acc1.Deposit(100.74)
	fmt.Println(acc1.GetBalance())
	acc1.Withdraw(80.34)
	fmt.Println(acc1.GetBalance())
	acc1.PrintInfo()

}
