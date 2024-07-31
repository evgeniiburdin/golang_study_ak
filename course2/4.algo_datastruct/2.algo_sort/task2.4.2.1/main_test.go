package main

import (
	"reflect"
	"testing"
	"time"
)

func TestByCount_Len(t *testing.T) {
	tests := []struct {
		name string
		p    ByCount
		want int
	}{
		{
			name: "testLen5",
			p:    ByCount{Product{}, Product{}, Product{}, Product{}, Product{}},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByCount_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		p    ByCount
		args args
		want bool
	}{
		{
			name: "testLess",
			p:    ByCount{Product{Count: 5}, Product{Count: 10}},
			args: args{
				i: 0,
				j: 1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByCount_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		p    ByCount
		args args
		want ByCount
	}{
		{
			name: "testSwap",
			p:    ByCount{Product{Count: 5}, Product{Count: 10}},
			args: args{
				i: 0,
				j: 1,
			},
			want: ByCount{Product{Count: 10}, Product{Count: 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.Swap(tt.args.i, tt.args.j)
			if !reflect.DeepEqual(tt.p, tt.want) {
				t.Errorf("Swap() = %v, want %v", tt.p, tt.want)
			}
		})
	}
}

func TestByCreatedAt_Len(t *testing.T) {
	tests := []struct {
		name string
		p    ByCreatedAt
		want int
	}{
		{
			name: "testLen5",
			p:    ByCreatedAt{Product{}, Product{}, Product{}, Product{}, Product{}},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByCreatedAt_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		p    ByCreatedAt
		args args
		want bool
	}{
		{
			name: "testLess",
			p:    ByCreatedAt{Product{CreatedAt: time.Now()}, Product{CreatedAt: time.Now().Add(5 * time.Second)}},
			args: args{
				i: 0,
				j: 1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByCreatedAt_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		p    ByCreatedAt
		args args
		want ByCreatedAt
	}{
		{
			name: "testSwap",
			p:    ByCreatedAt{Product{CreatedAt: time.Now()}, Product{CreatedAt: time.Now().Add(5 * time.Second)}},
			args: args{
				i: 0,
				j: 1,
			},
			want: ByCreatedAt{Product{CreatedAt: time.Now().Add(5 * time.Second)}, Product{CreatedAt: time.Now()}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.Swap(tt.args.i, tt.args.j)
		})
	}
}

func TestByPrice_Len(t *testing.T) {
	tests := []struct {
		name string
		p    ByPrice
		want int
	}{
		{
			name: "testLen5",
			p:    ByPrice{Product{}, Product{}, Product{}, Product{}, Product{}},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByPrice_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		p    ByPrice
		args args
		want bool
	}{
		{
			name: "testLess",
			p:    ByPrice{Product{Price: 4.99}, Product{Price: 10.99}},
			args: args{
				i: 0,
				j: 1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByPrice_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		p    ByPrice
		args args
		want ByPrice
	}{
		{
			name: "testSwap",
			p:    ByPrice{Product{Price: 4.99}, Product{Price: 10.99}},
			args: args{
				i: 0,
				j: 1,
			},
			want: ByPrice{Product{Price: 10.99}, Product{Price: 4.99}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.Swap(tt.args.i, tt.args.j)
		})
	}
}

func TestProduct_String(t *testing.T) {
	type fields struct {
		Name      string
		Price     float64
		CreatedAt time.Time
		Count     int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "testString",
			fields: fields{
				Name:      "test",
				Price:     4.99,
				CreatedAt: time.Now(),
				Count:     1,
			},
			want: "Name: test, Price: 4.99, Count: 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Product{
				Name:      tt.fields.Name,
				Price:     tt.fields.Price,
				CreatedAt: tt.fields.CreatedAt,
				Count:     tt.fields.Count,
			}
			if got := p.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
