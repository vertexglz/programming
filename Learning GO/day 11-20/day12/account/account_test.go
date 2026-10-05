package account

import (
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
				t.Errorf("unexpected error %v", err)
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
