package main

import (
	"net/http"
	"sync"
	"encoding/json"
	"strconv"
)

var (
	mu sync.Mutex
)

func getBooks(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	var list []Book
	for _, book := range books {
		list = append(list, book)
	}

	json.NewEncoder(w).Encode(list)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
	}

	book.ID = strconv.Itoa(len(books) + 1)

	mu.Lock()
	books[book.ID] = book
	mu.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func getBook(w http.ResponseWriter, r *http.Request, id string) {
	mu.Lock()
	book, ok := books[id]
	mu.Unlock()

	if !ok {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request, id string) {
	mu.Lock()
	_, ok := books[id]
	mu.Unlock()

	if !ok {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid Input", http.StatusBadRequest)
	}

	book.ID = id

	mu.Lock()
	books[book.ID] = book
	mu.Unlock()

	json.NewEncoder(w).Encode(book)
}

func deleteBook(w http.ResponseWriter, r *http.Request, id string) {
	mu.Lock()
	_, ok := books[id]
	if ok {
		delete(books, id)
	}
	mu.Unlock()

	if !ok {
		http.Error(w, "Not found", http.StatusNotFound)
	}

	w.WriteHeader(http.StatusNoContent)
}