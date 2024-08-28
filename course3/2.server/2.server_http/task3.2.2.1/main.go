package main

import (
	"net/http"
)

func main() {
	err := http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello World"))
		if err != nil {
			panic(err)
		}
	}))
	if err != nil {
		panic(err)
	}
}
