package account

import (
	"errors"
	"testing"
)

func TestDeposit(t *testing.T) {
	tests := []struct {
		name    string
		amount  float64
		wantErr bool
		wantBal float64
	}{
		{
			name:    "valid deposit",
			amount:  500,
			wantErr: false,
			wantBal: 1500,
		},
		{
			name:    "negative amount",
			amount:  -100,
			wantErr: true,
			wantBal: 1000,
		},
		{
			name:    "zero amount",
			amount:  0,
			wantErr: true,
			wantBal: 1000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			account, err := New("Pavel", 1000.00)

			if err != nil {
				t.Fatalf("unexpected error %v", err)
			}

			err = account.Deposit(tt.amount)

			if (err != nil) != tt.wantErr {
				t.Errorf("expected error: %v, got : %v", tt.wantErr, err)

			}
			if account.Balance() != tt.wantBal {
				t.Errorf("expected error:%v, got: %v", tt.wantBal, account.Balance())
			}

		})
	}
}

func TestWithdraw(t *testing.T) {
	tests := []struct {
		name    string
		amount  float64
		wantErr bool
		wantBal float64
	}{
		{name: "valid withdrawal",
			amount:  300,
			wantErr: false,
			wantBal: 700,
		},
		{name: "insufficient balance",
			amount:  1200,
			wantErr: true,
			wantBal: 1000,
		}, {
			name:    "zero amount",
			amount:  0,
			wantErr: true,
			wantBal: 1000,
		}, {
			name:    "negative amount",
			amount:  -100,
			wantErr: true,
			wantBal: 1000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			account, err := New("Pavel", 1000)

			if err != nil {
				t.Fatalf("unexpected error:%v", err)
			}

			err = account.Withdraw(tt.amount)

			if (err != nil) != tt.wantErr {
				t.Errorf("expected error: %v, got: %v", tt.wantErr, err)
			}
			if account.Balance() != tt.wantBal {
				t.Errorf("expected balance: %v , got:%v", tt.wantBal, account.Balance())
			}
		})
	}
}

func TestTransfer(t *testing.T) {
	tests := []struct {
		name        string
		amount      float64
		wantErr     bool
		wantFrom    float64
		wantTo      float64
		sameAccount bool
	}{
		{name: "valid transfer",
			amount:   300,
			wantErr:  false,
			wantFrom: 700,
			wantTo:   800},
		{name: "insufficient balance",
			amount:   1200,
			wantErr:  true,
			wantFrom: 1000,
			wantTo:   500},
		{name: "zero amount",
			amount:   0,
			wantErr:  true,
			wantFrom: 1000,
			wantTo:   500},
		{name: "negative amount",
			amount:   -100,
			wantErr:  true,
			wantFrom: 1000,
			wantTo:   500},
		{
			name:        "same account",
			amount:      100,
			wantErr:     true,
			wantFrom:    1000,
			wantTo:      1000,
			sameAccount: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			from, err := New("Pavel", 1000)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			var to *Account

			if tt.sameAccount {
				to = &from
			} else {
				account, err := New("Alex", 500)
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				to = &account
			}

			err = Transfer(&from, to, tt.amount)
			if (err != nil) != tt.wantErr {
				t.Errorf("expected error: %v, got: %v", tt.wantErr, err)
			}

			if from.Balance() != tt.wantFrom {
				t.Errorf("expected balance From %v, got %v", tt.wantFrom, from.Balance())
			}
			if to.Balance() != tt.wantTo {
				t.Errorf("expected balance To %v, got %v ", tt.wantTo, to.Balance())
			}
		})
	}

}

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		owner   string
		balance float64
		wantErr bool
	}{{
		name:    "valid account",
		owner:   "Pavel",
		balance: 1000.00,
		wantErr: false,
	}, {
		name:    "empty owner",
		owner:   "",
		balance: 1000.00,
		wantErr: true},
		{
			name:    "negative balance",
			owner:   "Pavel",
			balance: -100,
			wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			account, err := New(tt.owner, tt.balance)

			if (err != nil) != tt.wantErr {
				t.Errorf("expected error: %v, got: %v", tt.wantErr, err)
			}

			if tt.wantErr {
				return
			}

			if account.Owner() != tt.owner {
				t.Errorf("expected owner: %v, got: %v", tt.owner, account.Owner())
			}

			if account.Balance() != tt.balance {
				t.Errorf("expected balance: %v, got: %v", tt.balance, account.Balance())
			}
		})
	}
}

func TestWithdrawInsufficientBalance(t *testing.T) {
	account, err := New("Pavel", 1000.00)
	if err != nil {
		t.Fatalf("unexpected error : %v", err)
	}

	err = account.Withdraw(1500)

	if !errors.Is(err, ErrInsufficientBalance) {
		t.Errorf("expected balance: ErrInsufficientBalance , got : %v", err)
	}
}

func TestNewValidationError(t *testing.T) {
	_, err := New("", 1000.00)

	var validationErr ValidatorError

	if !errors.As(err, &validationErr) {
		t.Fatal("unexpected Validation Error")
	}
	if validationErr.Field != "owner" {
		t.Errorf("expected field owner, got %v", validationErr.Field)
	}

	if validationErr.Value != "empty" {
		t.Errorf("expected value empty, got %v", validationErr.Value)
	}
}

func TestNewNegativeBalanceError(t *testing.T) {
	_, err := New("Pavel", -100.00)
	var validationErr ValidatorError
	if !errors.As(err, &validationErr) {
		t.Fatal("unexpected Validation Error")
	}
	if validationErr.Field != "balance" {
		t.Errorf("expected field balance, got %v", validationErr.Field)
	}
	if validationErr.Value != "negative" {
		t.Errorf("expected value negative, got %v", validationErr.Value)
	}
}
