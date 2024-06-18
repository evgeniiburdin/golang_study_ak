package main

import (
	"bytes"
	"os"
	"reflect"
	"testing"
)

func TestMainFunc(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	_ = w.Close()
	os.Stdout = old

	var stdout bytes.Buffer
	_, _ = stdout.ReadFrom(r)

	expected := "Значение: 1, Новый срез: [2 3 4 5]"
	if stdout.String() != expected {
		t.Errorf("got %q, want %q", stdout.String(), expected)
	}
}

func TestPop(t *testing.T) {
	type args struct {
		xs []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 []int
	}{
		{"case1", args{[]int{1, 2, 3, 4, 5}}, 1, []int{2, 3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Pop(tt.args.xs)
			if got != tt.want {
				t.Errorf("Pop() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Pop() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
