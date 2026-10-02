package main

import (
	"errors"
	"fmt"
)

type Wallet struct {
	Balance float64
}

type Cash struct {
	Amount float64
}

type Card struct {
	Limit float64
}

type Payer interface {
	Pay(amount float64) error
}

func (w *Wallet) Pay(amount float64) error {
	if amount <= 0 {
		return errors.New("Некорректная сумма")
	} else if amount > w.Balance {
		return errors.New("Недостаточно средств")
	}
	w.Balance -= amount
	return nil
}

func (c *Card) Pay(amount float64) error {
	if amount <= 0 {
		return errors.New("Некорректная сумма")
	} else if amount > c.Limit {
		return errors.New("Недостаточно средств")
	}
	c.Limit -= amount
	return nil
}

func (c *Cash) Pay(amount float64) error {
	if amount <= 0 {
		return errors.New("Некорректная сумма")
	} else if amount > c.Amount {
		return errors.New("Недостаточно средств")
	}
	c.Amount -= amount
	return nil
}

func tryPayment(payers []Payer, amount float64) error {
	var lastErr error

	for _, value := range payers {
		err := value.Pay(amount)

		if err != nil {
			lastErr = err
			continue
		}

		return nil
	}

	return lastErr
}

func main() {
	payers := []Payer{
		&Wallet{Balance: 100},
		&Card{Limit: 500},
		&Cash{Amount: 50},
	}
	amount := 120.00

	fmt.Println(tryPayment(payers, amount))
}
