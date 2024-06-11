package mymath

import "testing"

func TestCeil(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"case1", args{4.5}, 5.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ceil(tt.args.x); got != tt.want {
				t.Errorf("Ceil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloor(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"case1", args{4.5}, 4.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Floor(tt.args.x); got != tt.want {
				t.Errorf("Floor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"case1", args{4.5, 8}, 8.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMin(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"case1", args{4.5, 8}, 4.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPow(t *testing.T) {
	type args struct {
		x float64
		n float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"case1", args{5.0, 2.0}, 25.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Pow(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("Pow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSqrt(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"case1", args{25.0}, 5.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sqrt(tt.args.x); got != tt.want {
				t.Errorf("Sqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}
