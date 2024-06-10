package main

import "testing"

func TestReplaceSymbols(t *testing.T) {
	type args struct {
		s   string
		old rune
		new rune
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case1", args{"Hello, world!", 'o', '0'}, "Hell0, w0rld!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceSymbols(tt.args.s, tt.args.old, tt.args.new); got != tt.want {
				t.Errorf("ReplaceSymbols() = %v, want %v", got, tt.want)
			}
		})
	}
}
