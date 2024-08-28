package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

var logger *zap.Logger

func main() {
	r := chi.NewRouter()

	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	r.Use(LoggerMiddleware)

	r.Get("/", handlerTest)
	r.Post("/post", handlerPost)
	r.Put("/put", handlerPut)
	r.Delete("/delete", handlerDelete)

	http.ListenAndServe(":8080", r)
}

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)

		logger.Info("Incoming request",
			zap.String("method", r.Method),
			zap.String("url", r.URL.String()),
			zap.String("remote_addr", r.RemoteAddr),
			zap.Duration("duration", duration),
		)
	})
}

func handlerTest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GET: Hello World"))
}

func handlerPost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("POST: Hello World"))
}

func handlerPut(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("PUT: Hello World"))
}

func handlerDelete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DELETE: Hello World"))
}
