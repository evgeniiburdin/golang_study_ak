package main

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestSQLiteGenerator_CreateTableSQL(t *testing.T) {
	sqlGenerator := &SQLiteGenerator{}
	user := User{}
	sql := sqlGenerator.CreateTableSQL(&user)
	expected := "CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, first_name VARCHAR(100), last_name VARCHAR(100), email VARCHAR(100) UNIQUE);"
	if sql != expected {
		t.Errorf("expected %v, got %v", expected, sql)
	}
}

func TestSQLiteGenerator_CreateInsertSQL(t *testing.T) {
	sqlGenerator := &SQLiteGenerator{}
	user := User{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
	}
	sql := sqlGenerator.CreateInsertSQL(&user)
	expected := "INSERT INTO users (id, first_name, last_name, email) VALUES ('1', 'John', 'Doe', 'john.doe@example.com');"
	if sql != expected {
		t.Errorf("expected %v, got %v", expected, sql)
	}
}

func TestMigrator_Migrate(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	sqlGenerator := &SQLiteGenerator{}
	migrator := NewMigrator(db, sqlGenerator)

	err = migrator.Migrate(&User{})
	if err != nil {
		t.Fatalf("migration failed: %v", err)
	}

	// Verify that the table was created
	var tableName string
	err = db.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='users';").Scan(&tableName)
	if err != nil {
		t.Fatalf("failed to verify table creation: %v", err)
	}
	if tableName != "users" {
		t.Errorf("expected table name 'users', got %v", tableName)
	}
}

func TestGoFakeitGenerator_GenerateFakeUser(t *testing.T) {
	fakeDataGenerator := &GoFakeitGenerator{}
	user := fakeDataGenerator.GenerateFakeUser()
	if user.FirstName == "" || user.LastName == "" || user.Email == "" {
		t.Errorf("expected non-empty fields, got %v", user)
	}
}
