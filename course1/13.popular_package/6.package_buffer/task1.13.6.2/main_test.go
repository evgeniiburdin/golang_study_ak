package main

import (
	"bytes"

	"testing"
)

func Test_getScanner(t *testing.T) {
	type args struct {
		b *bytes.Buffer
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case1", args{bytes.NewBuffer([]byte("Hello\n,\n World!"))}, "Hello, World!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := getScanner(tt.args.b)
			if scanner == nil {
				t.Errorf("scanner is nil")
			}
			result := ""
			for scanner.Scan() {
				result += scanner.Text()
			}
			if result != tt.want {
				t.Errorf("getScanner() got = %v, want %v", result, tt.want)
			}
		})
	}
}
