package main

import (
	"fmt"

	"strings"
)

type User struct {
	Name     string
	Comments []Comment
}

type Comment struct {
	Message string
}

func main() {
	users := []User{
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
	}

	users = FilterComments(users)
	fmt.Println(users)
}

func FilterComments(users []User) []User {
	var filteredUsers []User
	for _, user := range users {
		badComments := GetBadComments(user.Comments)
		var goodComments []Comment
		for _, comment := range user.Comments {
			if !isInBadComments(comment, badComments) {
				goodComments = append(goodComments, comment)
			}
		}
		user.Comments = goodComments
		filteredUsers = append(filteredUsers, user)
	}
	return filteredUsers
}

func IsBadComment(comment Comment) bool {
	return strings.Contains(strings.ToLower(comment.Message), "bad")
}

func GetBadComments(comments []Comment) []Comment {
	var badComments []Comment
	for _, comment := range comments {
		if IsBadComment(comment) {
			badComments = append(badComments, comment)
		}
	}
	return badComments
}

func isInBadComments(comment Comment, badComments []Comment) bool {
	for _, badComment := range badComments {
		if comment.Message == badComment.Message {
			return true
		}
	}
	return false
}
