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

func httpHandlerHelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

func httpHandlerHelloWorld2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world 2"))
}

func httpHandlerHelloWorld3(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world 3"))
}

func httpHandlerNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Not Found"))
}

func startHttpServer() {
	httpRouter := chi.NewRouter()

	httpRouter.Get("/", httpHandlerNotFound)
	httpRouter.Get("/1", httpHandlerHelloWorld)
	httpRouter.Get("/2", httpHandlerHelloWorld2)
	httpRouter.Get("/3", httpHandlerHelloWorld3)

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
