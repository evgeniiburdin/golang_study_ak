package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	jsoniter "github.com/json-iterator/go"
)

//go:generate easyjson -all user.go

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
}

// EasyJSON marshals the users using easyjson
func EasyJSON(users []User) ([]byte, error) {
	// github.com/mailru/easyjson seemed a very problematic library for me,
	// as it basically can not marshal a slice of structs into json, as well
	// as it can not marshal a single struct just by running "easyjson -all <file.go>".
	// After executing this command I've got a message "don't know how to decode int", which
	// references to the first "int" field of User struct. If we change it to the string type,
	// it starts giving a "don't know how to decode string" error
}

// JSON marshals the users using encoding/json
func JSON(users []User) ([]byte, error) {
	return json.Marshal(users)
}

// JSONiter marshals the users using jsoniter
func JSONiter(users []User) ([]byte, error) {
	jsoniterMarshaller := jsoniter.ConfigCompatibleWithStandardLibrary
	return jsoniterMarshaller.Marshal(users)
}

// GenerateUsers generates a list of users
func GenerateUsers(count int) []User {
	users := make([]User, count)
	for i := 0; i < count; i++ {
		users[i] = User{
			ID:       i,
			Username: gofakeit.Username(),
			Password: gofakeit.Password(true, true, true, true, false, 14),
			Age:      gofakeit.Number(18, 100),
			Email:    gofakeit.Email(),
		}
	}
	return users
}

// Benchmarks

func BenchmarkEasyJSON(b *testing.B) {
	users := GenerateUsers(1000)
	for i := 0; i < b.N; i++ {
		_, err := EasyJSON(users)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkJSON(b *testing.B) {
	users := GenerateUsers(1000)
	for i := 0; i < b.N; i++ {
		_, err := JSON(users)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkJSONiter(b *testing.B) {
	users := GenerateUsers(1000)
	for i := 0; i < b.N; i++ {
		_, err := JSONiter(users)
		if err != nil {
			b.Error(err)
		}
	}
}

func main() {
	users := GenerateUsers(10)
	data, err := EasyJSON(users)
	if err != nil {
		fmt.Println("Error marshaling with EasyJSON:", err)
	}
	fmt.Println("EasyJSON data:", string(data))

	data, err = JSON(users)
	if err != nil {
		fmt.Println("Error marshaling with JSON:", err)
	}
	fmt.Println("JSON data:", string(data))

	data, err = JSONiter(users)
	if err != nil {
		fmt.Println("Error marshaling with JSONiter:", err)
	}
	fmt.Println("JSONiter data:", string(data))
}
