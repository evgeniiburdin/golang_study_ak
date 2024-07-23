package main

import (
	"database/sql"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"library-server/internal/controller"
	"library-server/internal/repository"
	"library-server/internal/service"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	httpPort := os.Getenv("HTTP_LISTENER_PORT")

	connStr := "host=" + dbHost + " port=" + dbPort + " user=" + dbUser + " password=" + dbPass + " dbname=" + dbName + " sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo, _ := repository.NewPostgresRepository(db)
	svc := service.NewLibraryService(repo)
	ctrl := controller.NewLibraryController(svc)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Define routes
	r.Get("/authors", ctrl.ListAuthors)
	r.Get("/books", ctrl.ListBooks)
	r.Get("/users", ctrl.ListUsers)
	r.Post("/authors", ctrl.AddAuthor)
	r.Post("/books", ctrl.AddBook)
	r.Post("/users", ctrl.AddUser)
	r.Post("/rent", ctrl.RentBook)
	r.Post("/return", ctrl.ReturnBook)

	log.Printf("Server started on :%s", httpPort)
	log.Fatal(http.ListenAndServe(":"+httpPort, r))
}
