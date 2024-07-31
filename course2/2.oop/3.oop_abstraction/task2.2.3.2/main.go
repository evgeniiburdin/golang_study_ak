package main

import (
	"fmt"
)

// User struct defines the user model
type User struct {
	ID        int    `db_field:"id" db_type:"SERIAL PRIMARY KEY"`
	FirstName string `db_field:"first_name" db_type:"VARCHAR(100)"`
	LastName  string `db_field:"last_name" db_type:"VARCHAR(100)"`
	Email     string `db_field:"email" db_type:"VARCHAR(100) UNIQUE"`
}

// TableName returns the table name for the User model
func (u *User) TableName() string {
	return "users"
}

// Tabler interface defines a method to get the table name
type Tabler interface {
	TableName() string
}

// SQLGenerator interface defines methods for generating SQL queries
type SQLGenerator interface {
	CreateTableSQL(table Tabler) string
	CreateInsertSQL(model Tabler) string
}

// FakeDataGenerator interface defines methods for generating fake data
type FakeDataGenerator interface {
	GenerateFakeUser() User
}

// SQLiteGenerator struct implements the SQLGenerator interface
type SQLiteGenerator struct{}

// CreateTableSQL generates a SQL query to create a table
func (s *SQLiteGenerator) CreateTableSQL(table Tabler) string {
	panic("implement me")
}

// CreateInsertSQL generates a SQL query to insert data into a table
func (s *SQLiteGenerator) CreateInsertSQL(model Tabler) string {
	panic("implement me")
}

// GoFakeitGenerator struct implements the FakeDataGenerator interface
type GoFakeitGenerator struct{}

// GenerateFakeUser generates a fake User
func (g *GoFakeitGenerator) GenerateFakeUser() User {
	panic("implement me")
}

func main() {
	sqlGenerator := &SQLiteGenerator{}
	fakeDataGenerator := &GoFakeitGenerator{}
	user := User{}

	// Generate SQL to create the table
	sql := sqlGenerator.CreateTableSQL(&user)
	fmt.Println(sql)

	// Generate SQL to insert fake users
	for i := 0; i < 10; i++ {
		fakeUser := fakeDataGenerator.GenerateFakeUser()
		query := sqlGenerator.CreateInsertSQL(&fakeUser)
		fmt.Println(query)
	}
}
