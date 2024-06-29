package main

import "testing"

func TestGenerateRandomString(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name string
		args args
	}{
		{"case1", args{length: 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateRandomString(tt.args.length)
			if len(got) != tt.args.length {
				t.Errorf("GenerateRandomString() got len = %v, want %v", len(got), tt.args.length)
			}
			for _, b := range got {
				if !(b >= 'A' && b <= 'Z') {
					t.Errorf("GenerateRandomString() got = %v. Key must only contain of ASCII chars", got)
				}
			}
		})
	}
}
