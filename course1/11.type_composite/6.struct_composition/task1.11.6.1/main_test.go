package main

import (
	"bytes"
	"os"
	"reflect"
	"testing"
)

func TestOrder_AddDish(t *testing.T) {
	type fields struct {
		Dishes []Dish
		Total  float64
	}
	type args struct {
		dish Dish
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []Dish
	}{
		{"case1",
			fields{[]Dish{
				{"Burger", 5.90},
				{"Pizza", 10.90},
			}, 16.80},

			args{Dish{
				"Spaghetti", 2.90}},

			[]Dish{{
				"Burger", 5.90},
				{"Pizza", 10.90},
				{"Spaghetti", 2.90}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			order := &Order{
				Dishes: tt.fields.Dishes,
				Total:  tt.fields.Total,
			}
			order.AddDish(tt.args.dish)
			if got := order.Dishes; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddDish() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrder_CalculateTotal(t *testing.T) {
	type fields struct {
		Dishes []Dish
		Total  float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{"case1",
			fields{[]Dish{
				{"Burger", 5.90},
				{"Pizza", 10.90},
			}, 0},
			16.80},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			order := &Order{
				Dishes: tt.fields.Dishes,
				Total:  tt.fields.Total,
			}
			order.CalculateTotal()
			if got := order.Total; got != tt.want {
				t.Errorf("CalculateTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrder_RemoveDish(t *testing.T) {
	type fields struct {
		Dishes []Dish
		Total  float64
	}
	type args struct {
		dish Dish
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []Dish
	}{
		{"case1",
			fields{[]Dish{
				{"Burger", 5.90},
				{"Pizza", 10.90},
				{"Spaghetti", 2.90},
			}, 19.70},

			args{Dish{
				"Spaghetti", 2.90}},

			[]Dish{{
				"Burger", 5.90},
				{"Pizza", 10.90}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			order := &Order{
				Dishes: tt.fields.Dishes,
				Total:  tt.fields.Total,
			}
			order.RemoveDish(tt.args.dish)
			if got := order.Dishes; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveDish() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMainFunc(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	_ = w.Close()
	os.Stdout = old

	var stdout bytes.Buffer
	_, _ = stdout.ReadFrom(r)

	expected := "Total: 16.98\nTotal: 5.99\n"
	if stdout.String() != expected {
		t.Errorf("got %q, want %q", stdout.String(), expected)
	}
}
