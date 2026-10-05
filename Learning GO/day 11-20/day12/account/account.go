package account

import "errors"

type Account struct {
	balance float64
	owner   string
}

func New(owner string, balance float64) (Account, error) {
	if balance < 0 {
		return Account{}, errors.New("Некорректный баланс")
	} else if owner == "" {
		return Account{}, errors.New("Отсутствует имя владельца кошелька")
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
		return errors.New("Недостаточно баланса")
	}
	a.balance -= amount
	return nil
}

func Transfer(from *Account, to *Account, amount float64) error {
	if amount <= 0 {
		return errors.New("Некорректная сумма")
	}

	err := from.Withdraw(amount)
	if err != nil {
		return err
	}

	return to.Deposit(amount)
}
