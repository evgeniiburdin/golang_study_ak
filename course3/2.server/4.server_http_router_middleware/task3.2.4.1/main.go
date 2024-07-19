package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/r1", handleRoute1)
	r.Get("/r2", handleRoute2)
	r.Get("/r3", handleRoute3)

	http.ListenAndServe(":8080", r)
}

func handleRoute1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World! from route 1"))
}

func handleRoute2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World! from route 2"))
}

func handleRoute3(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World! from route 3"))
}
