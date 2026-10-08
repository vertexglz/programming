package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type User struct {
	Name  string `json:"user_name"`
	Age   int    `json:"user_age"`
	Email string `json:"email,omitempty"`
}

type Account struct {
	balance float64
	owner   string
}

func (a Account) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Owner   string  `json:"owner"`
		Balance float64 `json:"balance" `
	}{Owner: a.owner, Balance: a.balance})
}

func (a *Account) UnmarshalJSON(data []byte) error {
	var aux struct {
		Owner   string  `json:"owner"`
		Balance float64 `json:"balance"`
	}
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}

	if aux.Owner == "" {
		return errors.New("Owner is empty")
	}
	if aux.Balance < 0 {
		return errors.New("Balance is negative")
	}

	a.owner = aux.Owner
	a.balance = aux.Balance
	return nil

}

func main() {
	original := Account{
		balance: 750.50,
		owner:   "Alex",
	}

	data, err := json.Marshal(original)

	if err != nil {
		fmt.Println(err)
		return
	}

	var restored Account

	err = json.Unmarshal(data, &restored)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(restored.owner)
		fmt.Println(restored.balance)
	}

}
