package main

import "testing"

func Test_concatStrings(t *testing.T) {
	type args struct {
		xs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"emptyString", args{[]string{""}}, ""},
		{"singleString", args{[]string{"a"}}, "a"},
		{"multipleStrings", args{[]string{"a", "b", "c"}}, "abc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := concatStrings(tt.args.xs...); got != tt.want {
				t.Errorf("concatStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
