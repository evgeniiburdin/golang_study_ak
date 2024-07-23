package service

import (
	"library-server/internal/repository"
	"library-server/models"
)

type LibraryService interface {
	ListAuthors() ([]models.Author, error)
	ListBooks() ([]models.Book, error)
	ListUsers() ([]models.User, error)
	AddAuthor(name, biography string) error
	AddBook(title string, authorID int) error
	AddUser(name, email string) error
	RentBook(userID, bookID int) error
	ReturnBook(userID, bookID int) error
}

type LibraryServiceImpl struct {
	repo repository.Repository
}

func NewLibraryService(repo repository.Repository) *LibraryServiceImpl {
	return &LibraryServiceImpl{repo: repo}
}

func (s *LibraryServiceImpl) ListAuthors() ([]models.Author, error) {
	return s.repo.GetAuthors()
}

func (s *LibraryServiceImpl) ListBooks() ([]models.Book, error) {
	return s.repo.GetBooks()
}

func (s *LibraryServiceImpl) ListUsers() ([]models.User, error) {
	return s.repo.GetUsers()
}

func (s *LibraryServiceImpl) AddAuthor(name, biography string) error {
	return s.repo.AddAuthor(name, biography)
}

func (s *LibraryServiceImpl) AddBook(title string, authorID int) error {
	return s.repo.AddBook(title, authorID)
}

func (s *LibraryServiceImpl) AddUser(name, email string) error {
	return s.repo.AddUser(name, email)
}

func (s *LibraryServiceImpl) RentBook(userID, bookID int) error {
	return s.repo.RentBook(userID, bookID)
}

func (s *LibraryServiceImpl) ReturnBook(userID, bookID int) error {
	return s.repo.ReturnBook(userID, bookID)
}
