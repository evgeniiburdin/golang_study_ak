package models

type Book interface{}

type ConcreteBook struct {
	ID     int            `json:"id"`
	Title  string         `json:"title"`
	Author ConcreteAuthor `json:"author"`
}
