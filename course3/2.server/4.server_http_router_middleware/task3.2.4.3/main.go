package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

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

	r.Use(ProxyMiddleware)

	r.Get("/api/*", apiHandler)

	http.ListenAndServe(":8080", r)
}

func ProxyMiddleware(next http.Handler) http.Handler {
	hugoURL, _ := url.Parse("http://hugo:1313")
	proxy := httputil.NewSingleHostReverseProxy(hugoURL)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api/") {
			next.ServeHTTP(w, r)
		} else {
			proxy.ServeHTTP(w, r)
		}
	})
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from API"))
}
