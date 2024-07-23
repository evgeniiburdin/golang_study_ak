package models

type Author interface{}

type ConcreteAuthor struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Biography string `json:"biography"`
}
