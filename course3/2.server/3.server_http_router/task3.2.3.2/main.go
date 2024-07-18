package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	startHttpServer()
}

func startHttpServer() {
	httpRouter := chi.NewRouter()

	// Group 1
	httpRouter.Route("/group1", func(r chi.Router) {
		r.Get("/1", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Group 1 Привет, мир 1"))
		})
		r.Get("/2", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Group 1 Привет, мир 2"))
		})
		r.Get("/3", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Group 1 Привет, мир 3"))
		})
	})

	// Group 2
	httpRouter.Route("/group2", func(r chi.Router) {
		r.Get("/1", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Group 2 Привет, мир 1"))
		})
		r.Get("/2", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Group 2 Привет, мир 2"))
		})
		r.Get("/3", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Group 2 Привет, мир 3"))
		})
	})

	// Group 3
	httpRouter.Route("/group3", func(r chi.Router) {
		r.Get("/1", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Group 3 Привет, мир 1"))
		})
		r.Get("/2", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Group 3 Привет, мир 2"))
		})
		r.Get("/3", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Group 3 Привет, мир 3"))
		})
	})

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Running on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, httpRouter))
}
