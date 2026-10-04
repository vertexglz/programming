package wallet

import "errors"

type Wallet struct {
	balance float64
}

func New(balance float64) (Wallet, error) {
	if balance <= 0 {
		return Wallet{}, errors.New("Некорректное значение")
	}
	return Wallet{
		balance: balance,
	}, nil
}

func (w *Wallet) Balance() float64 {
	return w.balance
}
func (w *Wallet) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("Некорректное значение")
	}
	w.balance += amount
	return nil
}

func (w *Wallet) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("Некорректная сумма")
	} else if amount > w.balance {
		return errors.New("Недостаточный баланс")
	}
	w.balance -= amount
	return nil

}

func Transfer(from *Wallet, to *Wallet, amount float64) error {
	if amount <= 0 {
		return errors.New("Некорректная сумма")
	}

	err := from.Withdraw(amount)
	if err != nil {
		return err
	}

	return to.Deposit(amount)
}
