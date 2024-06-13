package main

import (
	"database/sql"
	"fmt"
	"log"

	squirrel "github.com/Masterminds/squirrel"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID       int
	Username string
	Email    string
}

func CreateUserTable() error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return err
	}
	defer db.Close()

	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL,
		email TEXT NOT NULL
	);`
	_, err = db.Exec(query)
	return err
}

func PrepareQuery(operation, table string, user User) (string, []interface{}, error) {
	var query squirrel.Sqlizer

	switch operation {
	case "insert":
		query = squirrel.Insert(table).Columns("username", "email").Values(user.Username, user.Email)
	case "select":
		query = squirrel.Select("id", "username", "email").From(table).Where(squirrel.Eq{"id": user.ID})
	case "update":
		query = squirrel.Update(table).SetMap(map[string]interface{}{
			"username": user.Username,
			"email":    user.Email,
		}).Where(squirrel.Eq{"id": user.ID})
	case "delete":
		query = squirrel.Delete(table).Where(squirrel.Eq{"id": user.ID})
	default:
		return "", nil, fmt.Errorf("unsupported operation: %s", operation)
	}

	sqlQuery, args, err := query.ToSql()
	return sqlQuery, args, err
}

func InsertUser(user User) error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return err
	}
	defer db.Close()

	query, args, err := PrepareQuery("insert", "users", user)
	if err != nil {
		return err
	}
	_, err = db.Exec(query, args...)
	return err
}

func SelectUser(userID int) (User, error) {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return User{}, err
	}
	defer db.Close()

	user := User{ID: userID}
	query, args, err := PrepareQuery("select", "users", user)
	if err != nil {
		return User{}, err
	}

	row := db.QueryRow(query, args...)
	err = row.Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return User{}, fmt.Errorf("user with id %d not found", userID)
		}
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

	query, args, err := PrepareQuery("update", "users", user)
	if err != nil {
		return err
	}
	_, err = db.Exec(query, args...)
	return err
}

func DeleteUser(userID int) error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return err
	}
	defer db.Close()

	user := User{ID: userID}
	query, args, err := PrepareQuery("delete", "users", user)
	if err != nil {
		return err
	}
	_, err = db.Exec(query, args...)
	return err
}

func main() {
	err := CreateUserTable()
	if err != nil {
		log.Fatalf("Error creating user table: %v", err)
	}

	user := User{Username: "JohnDoe", Email: "johndoe@example.com"}
	err = InsertUser(user)
	if err != nil {
		log.Fatalf("Error inserting user: %v", err)
	}

	retrievedUser, err := SelectUser(1)
	if err != nil {
		log.Fatalf("Error selecting user: %v", err)
	}
	fmt.Printf("Retrieved User: %+v\n", retrievedUser)

	retrievedUser.Email = "john.doe@newexample.com"
	err = UpdateUser(retrievedUser)
	if err != nil {
		log.Fatalf("Error updating user: %v", err)
	}

	err = DeleteUser(1)
	if err != nil {
		log.Fatalf("Error deleting user: %v", err)
	}
}
