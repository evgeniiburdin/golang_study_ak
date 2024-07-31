package main

import (
	"errors"
	"fmt"
)

// Accounter interface defines the methods for a bank account
type Accounter interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	Balance() float64
}

// CurrentAccount struct represents a current bank account
type CurrentAccount struct {
	balance float64
}

// Deposit method for CurrentAccount
func (a *CurrentAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be positive")
	}
	a.balance += amount
	return nil
}

// Withdraw method for CurrentAccount
func (a *CurrentAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("withdraw amount must be positive")
	}
	if amount > a.balance {
		return errors.New("insufficient funds")
	}
	a.balance -= amount
	return nil
}

// Balance method for CurrentAccount
func (a *CurrentAccount) Balance() float64 {
	return a.balance
}

// SavingsAccount struct represents a savings bank account
type SavingsAccount struct {
	balance float64
}

// Deposit method for SavingsAccount
func (a *SavingsAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be positive")
	}
	a.balance += amount
	return nil
}

// Withdraw method for SavingsAccount
func (a *SavingsAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("withdraw amount must be positive")
	}
	if amount > a.balance {
		return errors.New("insufficient funds")
	}
	if a.balance < 500 {
		return errors.New("cannot withdraw: balance less than 500")
	}
	a.balance -= amount
	return nil
}

// Balance method for SavingsAccount
func (a *SavingsAccount) Balance() float64 {
	return a.balance
}

// ProcessAccount function to perform deposit, withdraw, and print balance
func ProcessAccount(a Accounter) {
	a.Deposit(500)
	a.Withdraw(200)
	fmt.Printf("Balance: %.2f\n", a.Balance())
}

func main() {
	c := &CurrentAccount{}
	s := &SavingsAccount{}
	ProcessAccount(c)
	ProcessAccount(s)
}
