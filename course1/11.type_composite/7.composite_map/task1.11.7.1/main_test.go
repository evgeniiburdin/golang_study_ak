package task1_11_7_1

import (
	"reflect"
	"testing"
)

func Test_countWordOccurrences(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{"case1", args{"Lorem ipsum dolor sit amet consectetur adipiscing elit ipsum"},
			map[string]int{"lorem": 1, "ipsum": 2, "dolor": 1, "sit": 1, "amet": 1, "consectetur": 1, "adipiscing": 1, "elit": 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countWordOccurrences(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countWordOccurrences() = %v, want %v", got, tt.want)
			}
		})
	}
}
