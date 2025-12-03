# Simple CRUD REST API in Go

A minimal REST API implemented using Go’s standard `net/http` package.  
The service provides basic CRUD operations for managing an in-memory collection of books.

## Features

- Pure `net/http` implementation (no external frameworks)
- Full CRUD support:
  - `GET /books` – list all books
  - `POST /books` – create a new book
  - `GET /books/{id}` – retrieve a book by ID
  - `PUT /books/{id}` – update a book
  - `DELETE /books/{id}` – delete a book
- Thread-safe access to data using `sync.Mutex`
- Clear project structure split into separate files:
  - `main.go` — server startup
  - `routes.go` — request routing
  - `handlers.go` — CRUD logic

## Project Structure

.
├── handlers.go     # CRUD handlers
├── routes.go       # Routing for /books and /books/{id}
├── main.go         # Entry point and server setup
├── go.mod
└── go.sum
