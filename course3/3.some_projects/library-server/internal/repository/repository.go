package repository

import "library-server/models"

type Repository interface {
	GetAuthors() ([]models.Author, error)
	GetBooks() ([]models.Book, error)
	GetUsers() ([]models.User, error)
	AddAuthor(name string, biography string) error
	AddBook(title string, authorID int) error
	AddUser(name string, email string) error
	RentBook(userID int, bookID int) error
	ReturnBook(userID int, bookID int) error
}
