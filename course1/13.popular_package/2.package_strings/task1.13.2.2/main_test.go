package main

import (
	"reflect"
	"testing"
)

func TestFilterComments(t *testing.T) {
	type args struct {
		users []User
	}
	tests := []struct {
		name string
		args args
		want []User
	}{
		{"case1", args{[]User{
			{
				Name: "Betty",
				Comments: []Comment{
					{Message: "good Comment 1"},
					{Message: "BaD CoMmEnT 2"},
					{Message: "Bad Comment 3"},
					{Message: "Use camelCase please"},
				},
			},
			{
				Name: "Jhon",
				Comments: []Comment{
					{Message: "Good Comment 1"},
					{Message: "Good Comment 2"},
					{Message: "Good Comment 3"},
					{Message: "Bad Comments 4"},
				},
			},
		}}, []User{
			{
				Name: "Betty",
				Comments: []Comment{
					{Message: "good Comment 1"},
					{Message: "Use camelCase please"},
				},
			},
			{
				Name: "Jhon",
				Comments: []Comment{
					{Message: "Good Comment 1"},
					{Message: "Good Comment 2"},
					{Message: "Good Comment 3"},
				},
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterComments(tt.args.users); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterComments() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetBadComments(t *testing.T) {
	type args struct {
		comments []Comment
	}
	tests := []struct {
		name string
		args args
		want []Comment
	}{
		{"case1", args{[]Comment{Comment{"bad comment"}, Comment{"good comment"}}},
			[]Comment{Comment{"bad comment"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetBadComments(tt.args.comments); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBadComments() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsBadComment(t *testing.T) {
	type args struct {
		comment Comment
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1", args{Comment{"bAD comment"}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBadComment(tt.args.comment); got != tt.want {
				t.Errorf("IsBadComment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isInBadComments(t *testing.T) {
	type args struct {
		comment     Comment
		badComments []Comment
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1", args{Comment{"bad comment"}, []Comment{Comment{"bad comment"}}}, true},
		{"case1", args{Comment{"good comment"}, []Comment{Comment{"bad comment"}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isInBadComments(tt.args.comment, tt.args.badComments); got != tt.want {
				t.Errorf("isInBadComments() = %v, want %v", got, tt.want)
			}
		})
	}
}
