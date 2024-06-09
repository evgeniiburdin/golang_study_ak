package main

import (
	"bytes"
	"os"
	"reflect"
	"testing"
)

func TestCheckUpper(t *testing.T) {
	type args struct {
		old string
		new string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case2", args{"a", "a"}, "a"},
		{"case3", args{"A", "a"}, "A"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckUpper(tt.args.old, tt.args.new); got != tt.want {
				t.Errorf("CheckUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWordsToSentence(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case1", args{[]string{"w1", "w2"}}, "w1 w2!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WordsToSentence(tt.args.words); got != tt.want {
				t.Errorf("WordsToSentence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_filterSentence(t *testing.T) {
	type args struct {
		sentence  string
		censorMap map[string]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case1",
			args{"pass pass pass filter pass",
				map[string]string{"filter": "apple"}},
			"pass apple!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filterSentence(tt.args.sentence, tt.args.censorMap); got != tt.want {
				t.Errorf("filterSentence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_filterWords(t *testing.T) {
	type args struct {
		text      string
		censorMap map[string]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case1",
			args{"sentence sentence filter! Filter filter sentence!",
				map[string]string{"filter": "apple"}},
			"sentence apple! Apple sentence!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filterWords(tt.args.text, tt.args.censorMap); got != tt.want {
				t.Errorf("filterWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitSentences(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"case1", args{"abc, bnjk! bnjk, abc!"}, []string{"abc, bnjk", " bnjk, abc"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitSentences(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitSentences() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestMainFunc(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	_ = w.Close()
	os.Stdout = old

	var stdout bytes.Buffer
	_, _ = stdout.ReadFrom(r)

	expected := "Внимание! Покупай срочно фрукты только у нас! Яблоки по низким ценам! " +
		"Беги, успевай стать финансово независимым с помощью фруктов! Фрукты будущее финансового мира!\n"
	if stdout.String() != expected {
		t.Errorf("got %q, want %q", stdout.String(), expected)
	}
}
