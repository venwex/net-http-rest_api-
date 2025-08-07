package main

import (
	"net/http"
	"log"
	"fmt"
)

type Book struct {
	ID 		string `json:"id"`
	Title 	string `json:"title"`
	Author 	string `json:"author"`
}

var (
	books = make(map[string]Book)
)

func main() {
	http.HandleFunc("/books", booksHandler)
	http.HandleFunc("/books/", bookHandler)

	fmt.Println("Server is running on te 8080 port.")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
