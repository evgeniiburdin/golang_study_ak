package main

import (
	"errors"
	"sync"
)

// Account interface defines the methods for a bank account
type Account interface {
	Deposit(amount float64)
	Withdraw(amount float64) error
	Balance() float64
}

// SavingsAccount represents a savings account with a minimum balance restriction
type SavingsAccount struct {
	balance float64
	mu      sync.Mutex
}

// Deposit adds an amount to the SavingsAccount balance
func (a *SavingsAccount) Deposit(amount float64) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.balance += amount
}

// Withdraw removes an amount from the SavingsAccount balance if the balance after withdrawal is not less than 1000
func (a *SavingsAccount) Withdraw(amount float64) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.balance-amount < 1000 {
		return errors.New("cannot withdraw: balance cannot be less than 1000")
	}
	a.balance -= amount
	return nil
}

// Balance returns the current balance of the SavingsAccount
func (a *SavingsAccount) Balance() float64 {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.balance
}

// CheckingAccount represents a checking account without minimum balance restriction
type CheckingAccount struct {
	balance float64
	mu      sync.Mutex
}

// Deposit adds an amount to the CheckingAccount balance
func (a *CheckingAccount) Deposit(amount float64) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.balance += amount
}

// Withdraw removes an amount from the CheckingAccount balance if sufficient funds are available
func (a *CheckingAccount) Withdraw(amount float64) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.balance < amount {
		return errors.New("cannot withdraw: insufficient funds")
	}
	a.balance -= amount
	return nil
}

// Balance returns the current balance of the CheckingAccount
func (a *CheckingAccount) Balance() float64 {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.balance
}

// Customer represents a bank customer
type Customer struct {
	Name    string
	Account Account
}

// Option type for functional options
type Option func(*Customer)

// WithName sets the name of the customer
func WithName(name string) Option {
	return func(c *Customer) {
		c.Name = name
	}
}

// WithAccount sets the account of the customer
func WithAccount(account Account) Option {
	return func(c *Customer) {
		c.Account = account
	}
}

// NewCustomer creates a new Customer with the provided options
func NewCustomer(opts ...Option) *Customer {
	customer := &Customer{}
	for _, opt := range opts {
		opt(customer)
	}
	return customer
}
