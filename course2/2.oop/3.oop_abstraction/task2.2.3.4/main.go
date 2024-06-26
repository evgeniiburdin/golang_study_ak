package main

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"

	_ "github.com/mattn/go-sqlite3"
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

// Migrator struct handles the migration of database tables
type Migrator struct {
	db           *sql.DB
	sqlGenerator SQLGenerator
}

// NewMigrator creates a new instance of Migrator
func NewMigrator(db *sql.DB, sqlGenerator SQLGenerator) *Migrator {
	return &Migrator{db: db, sqlGenerator: sqlGenerator}
}

// Migrate creates tables for the given models
func (m *Migrator) Migrate(models ...Tabler) error {
	for _, model := range models {
		query := m.sqlGenerator.CreateTableSQL(model)
		_, err := m.db.Exec(query)
		if err != nil {
			return fmt.Errorf("failed to create table for model %v: %v", model.TableName(), err)
		}
	}
	return nil
}

// FakeDataGenerator interface defines methods for generating fake data
type FakeDataGenerator interface {
	GenerateFakeUser() User
}

// GoFakeitGenerator struct implements the FakeDataGenerator interface using gofakeit
type GoFakeitGenerator struct{}

// GenerateFakeUser generates a fake user
func (g *GoFakeitGenerator) GenerateFakeUser() User {
	return User{
		ID:        gofakeit.Number(1, 10000),
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Email:     gofakeit.Email(),
	}
}

// Main function
func main() {
	// Connect to SQLite database
	db, err := sql.Open("sqlite3", "file:my_database.db?cache=shared&mode=rwc")
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Create migrator with SQLiteGenerator
	sqlGenerator := &SQLiteGenerator{}
	migrator := NewMigrator(db, sqlGenerator)

	// Migrate User table
	if err := migrator.Migrate(&User{}); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}

	fmt.Println("Migration completed successfully")

	// Generate and print fake users
	fakeDataGenerator := &GoFakeitGenerator{}
	for i := 0; i < 10; i++ { // Adjust the number of fake users to generate
		fakeUser := fakeDataGenerator.GenerateFakeUser()
		query := sqlGenerator.CreateInsertSQL(&fakeUser)
		fmt.Println(query)
	}
}
