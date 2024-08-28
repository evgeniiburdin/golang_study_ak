package main

import (
	"testing"
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

func TestGoFakeitGenerator_GenerateFakeUser(t *testing.T) {
	fakeDataGenerator := &GoFakeitGenerator{}
	user := fakeDataGenerator.GenerateFakeUser()
	if user.FirstName == "" || user.LastName == "" || user.Email == "" {
		t.Errorf("expected non-empty fields, got %v", user)
	}
}

func TestGenerateUserInserts(t *testing.T) {
	sqlGenerator := &SQLiteGenerator{}
	fakeDataGenerator := &GoFakeitGenerator{}
	num := 10
	queries := GenerateUserInserts(fakeDataGenerator, sqlGenerator, num)
	if len(queries) != num {
		t.Errorf("expected %v queries, got %v", num, len(queries))
	}
	for _, query := range queries {
		if query == "" {
			t.Errorf("expected non-empty query, got empty string")
		}
	}
}
