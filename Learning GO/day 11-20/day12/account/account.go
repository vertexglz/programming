package account

import (
	"errors"
	"fmt"
)

var ErrInsufficientBalance = errors.New("недостаточно баланса")

type ValidatorError struct {
	Field string
	Value string
}
type Account struct {
	balance float64
	owner   string
}

func (e ValidatorError) Error() string {
	return fmt.Sprintf("invalid %s: %s", e.Field, e.Value)

}

func New(owner string, balance float64) (Account, error) {
	if balance < 0 {
		return Account{}, fmt.Errorf("не удалось создать аккаунт: %w",
			ValidatorError{Field: "balance",
				Value: "negative"})
	} else if owner == "" {
		return Account{}, fmt.Errorf("не удалось создать аккаунт : %w",
			ValidatorError{Field: "owner",
				Value: "empty"})
	}
	return Account{balance: balance, owner: owner}, nil
}

func (a Account) Owner() string {
	return a.owner
}

func (a Account) Balance() float64 {
	return a.balance
}

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("Некорректная сумма")
	}
	a.balance += amount
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("Некорректная сумма")
	} else if amount > a.balance {
		return fmt.Errorf("не удалось снять %.2f: %w", amount, ErrInsufficientBalance)
	}
	a.balance -= amount
	return nil
}

func Transfer(from *Account, to *Account, amount float64) error {
	if amount <= 0 {
		return errors.New("Некорректная сумма")
	}
	if from == to {
		return errors.New("Нельзя переводить самому себе")
	}

	err := from.Withdraw(amount)
	if err != nil {
		return err
	}

	return to.Deposit(amount)
}
