package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Comments []Comment `json:"comments"`
}

type Comment struct {
	Text string `json:"text"`
}

func getJSON(data []User) (string, error) {
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func main() {
	users := []User{
		{
			"Alice",
			35,
			[]Comment{
				{
					"comment1",
				},
				{
					"comment2",
				},
			},
		},
		{
			"Bob",
			56,
			[]Comment{
				{
					"comment3",
				},
				{
					"comment4",
				},
			},
		},
	}
	str, err := getJSON(users)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(str)
}
