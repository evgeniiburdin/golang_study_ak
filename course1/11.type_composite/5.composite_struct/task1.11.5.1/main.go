package main

import (
	"fmt"

	"github.com/icrowley/fake"
	"math/rand"
)

type User struct {
	Name string
	Age  int
}

func main() {
	users := getUsers()
	result := preparePrint(users)
	fmt.Println(result)
}

func getUsers() []User {
	usersAmount := 10
	users := make([]User, 0, usersAmount)
	for i := 0; i < usersAmount; i++ {
		users = append(users, User{
			fake.FirstName(),
			rand.Intn(100),
		})
	}
	return users
}

func preparePrint(users []User) string {
	str := ""
	for _, user := range users {
		str += fmt.Sprintf("Имя: %s, Возраст: %d\n", user.Name, user.Age)
	}
	return str
}
