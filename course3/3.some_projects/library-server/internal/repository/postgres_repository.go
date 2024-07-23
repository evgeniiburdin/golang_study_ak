package repository

import (
	"database/sql"
	"errors"
	"log"

	"github.com/brianvoe/gofakeit"

	"library-server/models"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) (Repository, error) {
	err := migrate(db)
	if err != nil {
		return nil, err
	}

	err = initializeWithData(db)
	if err != nil {
		return nil, err
	}

	return &PostgresRepository{
		db: db,
	}, nil
}

func migrate(db *sql.DB) error {
	tx, err := db.Begin()
	authorsTableCreationSQL :=
		`
CREATE TABLE IF NOT EXISTS authors (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    biography TEXT
);
`
	_, err = tx.Exec(authorsTableCreationSQL)
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			return err2
		}
		return err
	}

	booksTableCreationSQL :=
		`
CREATE TABLE IF NOT EXISTS books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author_id INT REFERENCES authors(id),
    UNIQUE(title, author_id)
);
`

	_, err = tx.Exec(booksTableCreationSQL)
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			return err2
		}
		return err
	}

	usersTableCreationSQL :=
		`
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash TEXT,
    UNIQUE(email)
);
`
	_, err = tx.Exec(usersTableCreationSQL)
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			return err2
		}
		return err
	}

	rentalsTableCreationSQL :=
		`
CREATE TABLE IF NOT EXISTS rentals (
    id SERIAL PRIMARY KEY,
    book_id INT REFERENCES books(id),
    user_id INT REFERENCES users(id),
    rented_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    returned_at TIMESTAMP,
    UNIQUE(book_id, user_id)
);
`
	_, err = tx.Exec(rentalsTableCreationSQL)
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			return err2
		}
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func initializeWithData(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	for i := 0; i < 10; i++ {
		name := gofakeit.Name()
		_, err := tx.Exec("INSERT INTO authors (name) VALUES ($1) ON CONFLICT (name) DO NOTHING", name)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}

	rows, err := tx.Query("SELECT id FROM authors")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer rows.Close()

	var authorIDs []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			log.Fatal(err)
			return err
		}
		authorIDs = append(authorIDs, id)
	}

	for i := 0; i < 100; i++ {
		title := gofakeit.Word()
		authorID := authorIDs[i%len(authorIDs)]
		_, err := tx.Exec("INSERT INTO books (title, author_id) VALUES ($1, $2) ON CONFLICT (title, author_id) DO NOTHING", title, authorID)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}

	for i := 0; i < 50; i++ {
		name := gofakeit.Name()
		email := gofakeit.Email()
		_, err := tx.Exec("INSERT INTO users (name, email) VALUES ($1, $2) ON CONFLICT (email) DO NOTHING", name, email)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *PostgresRepository) GetAuthors() ([]models.Author, error) {
	rows, err := r.db.Query("SELECT id, name, biography FROM authors")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var authors []models.Author
	for rows.Next() {
		var author models.ConcreteAuthor
		if err := rows.Scan(&author.ID, &author.Name, &author.Biography); err != nil {
			return nil, err
		}
		authors = append(authors, author)
	}
	return authors, nil
}

func (r *PostgresRepository) GetBooks() ([]models.Book, error) {
	rows, err := r.db.Query("SELECT b.id, b.title, a.id, a.name, a.biography FROM books b JOIN authors a ON b.author_id = a.id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.ConcreteBook
		if err := rows.Scan(&book.ID, &book.Title, &book.Author.ID, &book.Author.Name, &book.Author.Biography); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (r *PostgresRepository) GetUsers() ([]models.User, error) {
	rows, err := r.db.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.ConcreteUser
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *PostgresRepository) AddAuthor(name, biography string) error {
	_, err := r.db.Exec("INSERT INTO authors (name, biography) VALUES ($1, $2) ON CONFLICT (name) DO NOTHING", name, biography)
	return err
}

func (r *PostgresRepository) AddBook(title string, authorID int) error {
	_, err := r.db.Exec("INSERT INTO books (title, author_id) VALUES ($1, $2) ON CONFLICT (title, author_id) DO NOTHING", title, authorID)
	return err
}

func (r *PostgresRepository) AddUser(name, email string) error {
	_, err := r.db.Exec("INSERT INTO users (name, email) VALUES ($1, $2) ON CONFLICT (email) DO NOTHING", name, email)
	return err
}

func (r *PostgresRepository) RentBook(userID, bookID int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var bookCount int
	err = tx.QueryRow("SELECT COUNT(*) FROM rentals WHERE book_id = $1 AND returned_at IS NULL", bookID).Scan(&bookCount)
	if err != nil {
		return err
	}

	if bookCount > 0 {
		return errors.New("book is already rented")
	}

	_, err = tx.Exec("INSERT INTO rentals (book_id, user_id) VALUES ($1, $2)", bookID, userID)
	return err
}

func (r *PostgresRepository) ReturnBook(userID, bookID int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var rentalCount int
	err = tx.QueryRow("SELECT COUNT(*) FROM rentals WHERE book_id = $1 AND user_id = $2 AND returned_at IS NULL", bookID, userID).Scan(&rentalCount)
	if err != nil {
		return err
	}

	if rentalCount == 0 {
		return errors.New("no active rental found for this book and user")
	}

	_, err = tx.Exec("UPDATE rentals SET returned_at = CURRENT_TIMESTAMP WHERE book_id = $1 AND user_id = $2 AND returned_at IS NULL", bookID, userID)
	return err
}
