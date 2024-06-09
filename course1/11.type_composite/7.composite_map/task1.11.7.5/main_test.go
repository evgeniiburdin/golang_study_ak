package main

import "testing"

func Test_filterSentence(t *testing.T) {
	type args struct {
		sentence string
		filter   map[string]bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case1", args{"Lorem ipsum dolor sit amet consectetur adipiscing elit ipsum",
			map[string]bool{"ipsum": true, "elit": true}},
			"Lorem dolor sit amet consectetur adipiscing"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filterSentence(tt.args.sentence, tt.args.filter); got != tt.want {
				t.Errorf("filterSentence() = %v, want %v", got, tt.want)
			}
		})
	}
}
