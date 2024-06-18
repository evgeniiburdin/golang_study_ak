package main

import (
	"encoding/json"

	"fmt"
)

type User struct {
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Comments []Comment `json:"comments"`
}

type Comment struct {
	Text string `json:"text"`
}

func getUsersFromJSON(data []byte) ([]User, error) {
	var users []User
	err := json.Unmarshal(data, &users)
	return users, err
}

func main() {
	data := []byte(`
		
			[
				{
					"name": "Alice",
					"age": 35,
					"comments": [
						{
							"text": "comment1"
						},
						{
							"text": "comment2"
						}
					]
				},
				{
					"name": "Bob",
					"age": 40,
					"comments": [
						{
							"text": "comment3"
						},
						{
							"text": "comment4"
						}
					]
				}
			]
		
`)
	users, err := getUsersFromJSON(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, user := range users {
		fmt.Println("Name:", user.Name)
		fmt.Println("Age:", user.Age)
		fmt.Println("Comments:")
		for _, comment := range user.Comments {
			fmt.Println("Comment:", comment.Text)
		}
		fmt.Println()
	}
}
