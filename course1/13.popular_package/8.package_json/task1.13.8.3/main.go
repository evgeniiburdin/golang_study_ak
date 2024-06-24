package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type User struct {
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Comments []Comment `json:"comments"`
}

type Comment struct {
	Text string `json:"text"`
}

func writeJSON(filePath string, data []User) error {
	writeData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	dir := filepath.Dir(filePath)
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			fmt.Printf("Error closing file %s: %s\n", filePath, cerr)
		}
	}()

	_, err = file.Write(writeData)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	data := []User{
		User{
			Name: "Alice",
			Age:  35,
			Comments: []Comment{
				Comment{
					Text: "comment1",
				},
				Comment{
					Text: "comment2",
				},
			},
		},
		User{
			Name: "Bob",
			Age:  56,
			Comments: []Comment{
				Comment{
					Text: "comment3",
				},
				Comment{
					Text: "comment4",
				},
			},
		},
	}
	err := writeJSON("C:\\Users\\Eug\\Desktop\\text1.txt", data)
	if err != nil {
		fmt.Printf("Error writing to file %s\n", err)
	}
}
