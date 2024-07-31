package main

import (
	"errors"
	"fmt"
)

// Order interface defines the methods for managing an order
type Order interface {
	AddItem(item string, quantity int) error
	RemoveItem(item string) error
	GetOrderDetails() map[string]int
}

// DineInOrder struct represents an order for dining in
type DineInOrder struct {
	orderDetails map[string]int
}

// AddItem method for DineInOrder
func (d *DineInOrder) AddItem(item string, quantity int) error {
	if quantity <= 0 {
		return errors.New("quantity must be greater than 0")
	}
	if _, exists := d.orderDetails[item]; exists {
		d.orderDetails[item] += quantity
	} else {
		d.orderDetails[item] = quantity
	}
	return nil
}

// RemoveItem method for DineInOrder
func (d *DineInOrder) RemoveItem(item string) error {
	if _, exists := d.orderDetails[item]; exists {
		delete(d.orderDetails, item)
		return nil
	}
	return errors.New("item does not exist in the order")
}

// GetOrderDetails method for DineInOrder
func (d *DineInOrder) GetOrderDetails() map[string]int {
	return d.orderDetails
}

// TakeAwayOrder struct represents an order for takeaway
type TakeAwayOrder struct {
	orderDetails map[string]int
}

// AddItem method for TakeAwayOrder
func (t *TakeAwayOrder) AddItem(item string, quantity int) error {
	if quantity <= 0 {
		return errors.New("quantity must be greater than 0")
	}
	if _, exists := t.orderDetails[item]; exists {
		t.orderDetails[item] += quantity
	} else {
		t.orderDetails[item] = quantity
	}
	return nil
}

// RemoveItem method for TakeAwayOrder
func (t *TakeAwayOrder) RemoveItem(item string) error {
	if _, exists := t.orderDetails[item]; exists {
		delete(t.orderDetails, item)
		return nil
	}
	return errors.New("item does not exist in the order")
}

// GetOrderDetails method for TakeAwayOrder
func (t *TakeAwayOrder) GetOrderDetails() map[string]int {
	return t.orderDetails
}

// ManageOrder function to manage the order by adding, removing items and printing the order details
func ManageOrder(o Order) {
	o.AddItem("Pizza", 2)
	o.AddItem("Burger", 1)
	o.RemoveItem("Pizza")
	fmt.Println(o.GetOrderDetails())
}

func main() {
	dineIn := &DineInOrder{orderDetails: make(map[string]int)}
	takeAway := &TakeAwayOrder{orderDetails: make(map[string]int)}

	fmt.Println("DineIn Order:")
	ManageOrder(dineIn)

	fmt.Println("TakeAway Order:")
	ManageOrder(takeAway)
}
