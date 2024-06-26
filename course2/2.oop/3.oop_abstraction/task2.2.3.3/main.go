package main

import (
	"fmt"
	"reflect"
	"strings"
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
	t := reflect.TypeOf(table).Elem()
	var columns []string
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		column := fmt.Sprintf("%s %s", field.Tag.Get("db_field"), field.Tag.Get("db_type"))
		columns = append(columns, column)
	}
	return fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s);", table.TableName(), strings.Join(columns, ", "))
}

// CreateInsertSQL generates a SQL query to insert data into a table
func (s *SQLiteGenerator) CreateInsertSQL(model Tabler) string {
	v := reflect.ValueOf(model).Elem()
	t := reflect.TypeOf(model).Elem()

	var columns, placeholders []string
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		columns = append(columns, field.Tag.Get("db_field"))
		placeholders = append(placeholders, fmt.Sprintf("'%v'", v.Field(i).Interface()))
	}
	return fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);", model.TableName(), strings.Join(columns, ", "), strings.Join(placeholders, ", "))
}

// GoFakeitGenerator struct implements the FakeDataGenerator interface
type GoFakeitGenerator struct{}

// GenerateFakeUser generates a fake User
func (g *GoFakeitGenerator) GenerateFakeUser() User {
	return User{
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Email:     gofakeit.Email(),
	}
}

// GenerateUserInserts generates a specified number of insert queries
func GenerateUserInserts(generator FakeDataGenerator, sqlGenerator SQLGenerator, num int) []string {
	var queries []string
	for i := 0; i < num; i++ {
		fakeUser := generator.GenerateFakeUser()
		query := sqlGenerator.CreateInsertSQL(&fakeUser)
		queries = append(queries, query)
	}
	return queries
}

func main() {
	sqlGenerator := &SQLiteGenerator{}
	fakeDataGenerator := &GoFakeitGenerator{}
	user := User{}

	// Generate SQL to create the table
	sql := sqlGenerator.CreateTableSQL(&user)
	fmt.Println(sql)

	// Generate SQL to insert fake users
	for i := 0; i < 34; i++ {
		fakeUser := fakeDataGenerator.GenerateFakeUser()
		query := sqlGenerator.CreateInsertSQL(&fakeUser)
		fmt.Println(query)
	}

	// Generate multiple user insert queries
	queries := GenerateUserInserts(fakeDataGenerator, sqlGenerator, 34)
	for _, query := range queries {
		fmt.Println(query)
	}
}
