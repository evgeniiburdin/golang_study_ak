package main

import (
	"fmt"
	"github.com/google/btree"
)

// User represents a user with an ID, Name, and Age
type User struct {
	ID   int
	Name string
	Age  int
}

// Less implements the Less method of the btree.Item interface for User
func (u User) Less(than btree.Item) bool {
	return u.ID < than.(User).ID
}

// BTree represents a B-tree
type BTree struct {
	tree *btree.BTree
}

// NewBTree creates a new BTree with the given degree
func NewBTree(degree int) *BTree {
	return &BTree{tree: btree.New(degree)}
}

// Insert inserts a user into the BTree
func (bt *BTree) Insert(user User) {
	bt.tree.ReplaceOrInsert(user)
}

// Search searches for a user with the given ID in the BTree
func (bt *BTree) Search(id int) *User {
	item := bt.tree.Get(User{ID: id})
	if item == nil {
		return nil
	}
	user := item.(User)
	return &user
}

func main() {
	bt := NewBTree(2)
	users := []User{
		{1, "Alice", 30},
		{2, "Bob", 25},
		{3, "Charlie", 35},
		// добавьте больше пользователей при необходимости
	}

	for _, user := range users {
		bt.Insert(user)
	}

	if user := bt.Search(2); user != nil {
		fmt.Printf("Найден пользователь: %v\n", *user)
	} else {
		fmt.Println("Пользователь не найден")
	}
}
