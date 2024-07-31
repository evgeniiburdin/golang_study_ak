package main

import (
	"testing"
)

func TestSQLiteGenerator_CreateTableSQL(t *testing.T) {
	sqlGenerator := &SQLiteGenerator{}
	defer func() {
		if r := recover(); r != nil {
			if r != "implement me" {
				t.Errorf("expected panic with 'implement me', got %v", r)
			}
		} else {
			t.Errorf("expected panic, got none")
		}
	}()
	sqlGenerator.CreateTableSQL(&User{})
}

func TestSQLiteGenerator_CreateInsertSQL(t *testing.T) {
	sqlGenerator := &SQLiteGenerator{}
	defer func() {
		if r := recover(); r != nil {
			if r != "implement me" {
				t.Errorf("expected panic with 'implement me', got %v", r)
			}
		} else {
			t.Errorf("expected panic, got none")
		}
	}()
	sqlGenerator.CreateInsertSQL(&User{})
}

func TestGoFakeitGenerator_GenerateFakeUser(t *testing.T) {
	fakeDataGenerator := &GoFakeitGenerator{}
	defer func() {
		if r := recover(); r != nil {
			if r != "implement me" {
				t.Errorf("expected panic with 'implement me', got %v", r)
			}
		} else {
			t.Errorf("expected panic, got none")
		}
	}()
	fakeDataGenerator.GenerateFakeUser()
}
