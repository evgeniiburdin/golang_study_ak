package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func CreateUserTable() error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return err
	}
	defer db.Close()

	stmt := `CREATE TABLE IF NOT EXISTS users (
    		id INTEGER PRIMARY KEY AUTOINCREMENT,
    		name TEXT NOT NULL,
    		age INTEGER NOT NULL
			);`
	_, err = db.Exec(stmt)

	return err
}

func InsertUser(user User) error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return err
	}
	defer db.Close()

	stmt := `INSERT INTO users (name, age) VALUES (?, ?)`
	_, err = db.Exec(stmt, user.Name, user.Age)

	return err
}

func SelectUser(id int) (User, error) {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return User{}, err
	}

	defer db.Close()

	stmt := `SELECT * FROM users WHERE id = ?`
	row := db.QueryRow(stmt, id)
	var user User
	err = row.Scan(&user.ID, &user.Name, &user.Age)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func UpdateUser(user User) error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return err
	}

	defer db.Close()

	stmt := `UPDATE users SET name = ?, age = ? WHERE id = ?`
	_, err = db.Exec(stmt, user.Name, user.Age, user.ID)

	return err
}

func DeleteUser(id int) error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return err
	}

	defer db.Close()

	stmt := `DELETE FROM users WHERE id = ?`
	_, err = db.Exec(stmt, id)

	return err
}

func main() {
	err := CreateUserTable()
	if err != nil {
		log.Fatalf("Error creating user table: %v", err)
	}

	user := User{Name: "John Doe", Age: 30}
	err = InsertUser(user)
	if err != nil {
		log.Fatalf("Error inserting user: %v", err)
	}

	retrievedUser, err := SelectUser(1)
	if err != nil {
		log.Fatalf("Error selecting user: %v", err)
	}
	fmt.Printf("Retrieved User: %+v\n", retrievedUser)

	retrievedUser.Age = 31
	err = UpdateUser(retrievedUser)
	if err != nil {
		log.Fatalf("Error updating user: %v", err)
	}

	err = DeleteUser(1)
	if err != nil {
		log.Fatalf("Error deleting user: %v", err)
	}
}
