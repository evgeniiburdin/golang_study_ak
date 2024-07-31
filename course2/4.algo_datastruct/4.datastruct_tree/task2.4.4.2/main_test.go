package main

import (
	"testing"
)

// TestNewBTree tests the creation of a new BTree
func TestNewBTree(t *testing.T) {
	bt := NewBTree(2)
	if bt.tree == nil {
		t.Errorf("Expected BTree to be initialized, but got nil")
	}
}

// TestInsertAndSearch tests the insertion and searching of users in the BTree
func TestInsertAndSearch(t *testing.T) {
	bt := NewBTree(2)
	users := []User{
		{1, "Alice", 30},
		{2, "Bob", 25},
		{3, "Charlie", 35},
	}

	for _, user := range users {
		bt.Insert(user)
	}

	tests := []struct {
		id       int
		expected *User
	}{
		{1, &User{1, "Alice", 30}},
		{2, &User{2, "Bob", 25}},
		{3, &User{3, "Charlie", 35}},
		{4, nil},
	}

	for _, test := range tests {
		result := bt.Search(test.id)
		if result == nil && test.expected != nil {
			t.Errorf("Expected to find user %v, but got nil", test.expected)
		} else if result != nil && test.expected == nil {
			t.Errorf("Expected to find no user, but got %v", result)
		} else if result != nil && *result != *test.expected {
			t.Errorf("Expected to find user %v, but got %v", test.expected, result)
		}
	}
}

func TestUserLess(t *testing.T) {
	u1 := User{ID: 1}
	u2 := User{ID: 2}
	if !u1.Less(u2) {
		t.Errorf("Expected %v to be less than %v", u1, u2)
	}
	if u2.Less(u1) {
		t.Errorf("Expected %v to not be less than %v", u2, u1)
	}
}
