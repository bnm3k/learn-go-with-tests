package c7

import (
	"errors"
	"fmt"
)

// Bitcoin readme todo
type Bitcoin int

//String Stringer
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Stringer readme todo
type Stringer interface {
	String() string
}

// Wallet readme todo
type Wallet struct {
	balance Bitcoin
}

// Deposit readme todo
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance readme todo
func (w *Wallet) Balance() Bitcoin { return w.balance }

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
