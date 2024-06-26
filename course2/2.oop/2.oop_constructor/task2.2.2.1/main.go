package main

import (
	"time"
)

type Order struct {
	ID         int
	CustomerID string
	Items      []string
	OrderDate  time.Time
}

type Option func(*Order)

func WithCustomerID(CustomerID string) Option {
	return func(o *Order) {
		o.CustomerID = CustomerID
	}
}

func WithItems(Items []string) Option {
	return func(o *Order) {
		o.Items = Items
	}
}

func WithOrderDate(OrderDate time.Time) Option {
	return func(o *Order) {
		o.OrderDate = OrderDate
	}
}

func NewOrder(ID int, options ...Option) *Order {
	order := &Order{
		ID: ID,
	}
	for _, option := range options {
		option(order)
	}
	return order
}
