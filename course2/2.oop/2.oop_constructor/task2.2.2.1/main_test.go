package main

import (
	"reflect"
	"testing"
	"time"
)

func TestNewOrder(t *testing.T) {
	type args struct {
		ID      int
		options []Option
	}
	tests := []struct {
		name string
		args args
		want *Order
	}{
		{
			name: "test1",
			args: args{
				ID: 3,
				options: []Option{
					WithCustomerID("123"),
					WithItems([]string{"item1", "item2"}),
					WithOrderDate(time.Date(2024, time.June, 26, 12, 0, 0, 0, time.UTC)),
				},
			},
			want: &Order{3, "123", []string{"item1", "item2"}, time.Date(2024, time.June, 26, 12, 0, 0, 0, time.UTC)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOrder(tt.args.ID, tt.args.options...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithCustomerID(t *testing.T) {
	type args struct {
		CustomerID string
	}
	tests := []struct {
		name string
		args args
		want Order
	}{
		{
			name: "test1",
			args: args{
				CustomerID: "123",
			},
			want: Order{ID: 3, CustomerID: "123"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			option := WithCustomerID(tt.args.CustomerID)
			order := *NewOrder(3, option)
			if !reflect.DeepEqual(order, tt.want) {
				t.Errorf("WithCustomerID() = %v, want %v", order, tt.want)
			}
		})
	}
}

func TestWithItems(t *testing.T) {
	type args struct {
		Items []string
	}
	tests := []struct {
		name string
		args args
		want Order
	}{
		{
			name: "test1",
			args: args{
				Items: []string{"item1", "item2"},
			},
			want: Order{ID: 3, Items: []string{"item1", "item2"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			option := WithItems(tt.args.Items)
			order := *NewOrder(3, option)
			if !reflect.DeepEqual(order, tt.want) {
				t.Errorf("WithItems() = %v, want %v", order, tt.want)
			}
		})
	}
}

func TestWithOrderDate(t *testing.T) {
	type args struct {
		OrderDate time.Time
	}
	tests := []struct {
		name string
		args args
		want Order
	}{
		{
			name: "test1",
			args: args{
				OrderDate: time.Date(2024, time.June, 26, 12, 0, 0, 0, time.UTC),
			},
			want: Order{ID: 3, OrderDate: time.Date(2024, time.June, 26, 12, 0, 0, 0, time.UTC)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			option := WithOrderDate(tt.args.OrderDate)
			order := *NewOrder(3, option)
			if !reflect.DeepEqual(order, tt.want) {
				t.Errorf("WithOrderDate() = %v, want %v", order, tt.want)
			}
		})
	}
}
