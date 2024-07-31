package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type User struct {
	Name     string    `yaml:"name"`
	Age      int       `yaml:"age"`
	Comments []Comment `yaml:"comments"`
}

type Comment struct {
	Text string `yaml:"text"`
}

func writeYAML(filePath string, data interface{}) error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		dir := filepath.Dir(filePath)
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	file, err := os.Create(filepath.Clean(filePath))
	if err != nil {
		return err
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			fmt.Printf("Error closing file: %v", cerr)
		}
	}()

	dataYAML, err := yaml.Marshal(data)
	if err != nil {
		return err
	}

	_, err = file.Write(dataYAML)
	return err
}

func main() {
	users := []User{
		{
			Name: "Jane",
			Age:  31,
			Comments: []Comment{
				{
					"comment1",
				},
				{
					"comment2",
				},
			},
		},
		{
			Name: "Joe",
			Age:  32,
			Comments: []Comment{
				{
					"filepath",
				},
				{
					"filepath2",
				},
			},
		},
	}
	err := writeYAML("./users.yaml", users)
	if err != nil {
		fmt.Println(err)
	}
}
