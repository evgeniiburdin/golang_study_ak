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

func writeJSON(filePath string, data interface{}) error {
	writeData, err := json.MarshalIndent(data, "", "  ")
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
	err := writeJSON("C:\\Users\\Eug\\Desktop\\text2.txt", data)
	if err != nil {
		fmt.Printf("Error writing to file %s\n", err)
	}
}
