package controller

import (
	"encoding/json"
	"net/http"

	"library-server/internal/service"
)

type LibraryController struct {
	svc service.LibraryService
}

func NewLibraryController(svc service.LibraryService) *LibraryController {
	return &LibraryController{svc: svc}
}

func (c *LibraryController) ListAuthors(w http.ResponseWriter, r *http.Request) {
	authors, err := c.svc.ListAuthors()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(authors)
}

func (c *LibraryController) ListBooks(w http.ResponseWriter, r *http.Request) {
	books, err := c.svc.ListBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(books)
}

func (c *LibraryController) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := c.svc.ListUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func (c *LibraryController) AddAuthor(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name      string `json:"name"`
		Biography string `json:"biography"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := c.svc.AddAuthor(input.Name, input.Biography); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (c *LibraryController) AddBook(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title    string `json:"title"`
		AuthorID int    `json:"author_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := c.svc.AddBook(input.Title, input.AuthorID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (c *LibraryController) AddUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := c.svc.AddUser(input.Name, input.Email); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (c *LibraryController) RentBook(w http.ResponseWriter, r *http.Request) {
	var input struct {
		UserID int `json:"user_id"`
		BookID int `json:"book_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := c.svc.RentBook(input.UserID, input.BookID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *LibraryController) ReturnBook(w http.ResponseWriter, r *http.Request) {
	var input struct {
		UserID int `json:"user_id"`
		BookID int `json:"book_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := c.svc.ReturnBook(input.UserID, input.BookID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
