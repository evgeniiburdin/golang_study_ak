package main

import (
	"github.com/go-chi/chi"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGroup1Routes(t *testing.T) {
	router := chi.NewRouter()
	router.Route("/group1", func(r chi.Router) {
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

	tests := []struct {
		route    string
		expected string
	}{
		{"/group1/1", "Group 1 Привет, мир 1"},
		{"/group1/2", "Group 1 Привет, мир 2"},
		{"/group1/3", "Group 1 Привет, мир 3"},
	}

	for _, test := range tests {
		req, err := http.NewRequest("GET", test.route, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		if rr.Body.String() != test.expected {
			t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), test.expected)
		}
	}
}

func TestGroup2Routes(t *testing.T) {
	router := chi.NewRouter()
	router.Route("/group2", func(r chi.Router) {
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

	tests := []struct {
		route    string
		expected string
	}{
		{"/group2/1", "Group 2 Привет, мир 1"},
		{"/group2/2", "Group 2 Привет, мир 2"},
		{"/group2/3", "Group 2 Привет, мир 3"},
	}

	for _, test := range tests {
		req, err := http.NewRequest("GET", test.route, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		if rr.Body.String() != test.expected {
			t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), test.expected)
		}
	}
}

func TestGroup3Routes(t *testing.T) {
	router := chi.NewRouter()
	router.Route("/group3", func(r chi.Router) {
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

	tests := []struct {
		route    string
		expected string
	}{
		{"/group3/1", "Group 3 Привет, мир 1"},
		{"/group3/2", "Group 3 Привет, мир 2"},
		{"/group3/3", "Group 3 Привет, мир 3"},
	}

	for _, test := range tests {
		req, err := http.NewRequest("GET", test.route, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		if rr.Body.String() != test.expected {
			t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), test.expected)
		}
	}
}
