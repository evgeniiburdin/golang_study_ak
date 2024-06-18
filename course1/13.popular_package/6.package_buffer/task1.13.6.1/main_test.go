package main

import (
	"bytes"
	
	"testing"
)

func Test_getReader(t *testing.T) {
	type args struct {
		b *bytes.Buffer
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case1", args{bytes.NewBufferString("Hello, World!")}, "Hello, World!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := getReader(tt.args.b)
			b := make([]byte, 13)
			_, _ = reader.Read(b)
			if string(b) != tt.want {
				t.Errorf("getReader() got = %v, want %v", string(b), tt.want)
			}
		})
	}
}
