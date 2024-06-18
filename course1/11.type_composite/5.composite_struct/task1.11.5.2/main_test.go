package main

import (
	"reflect"
	"testing"
)

func Test_getAnimals(t *testing.T) {
	if got := getAnimals(); len(got) != 3 || reflect.TypeOf(got) != reflect.TypeOf([]Animal{}) {
		t.Errorf("getAnimals() = %v, want slice of Animal{} of len 3", got)
	}
}

func Test_preparePrint(t *testing.T) {
	type args struct {
		animals []Animal
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case1", args{[]Animal{{"animal", "Paul", 7}, {"animal", "Rose", 5}}},
			"Тип: animal, Имя: Paul, Возраст: 7\nТип: animal, Имя: Rose, Возраст: 5\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preparePrint(tt.args.animals); got != tt.want {
				t.Errorf("preparePrint() = %v, want %v", got, tt.want)
			}
		})
	}
}
