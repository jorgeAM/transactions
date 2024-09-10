package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	db, err := sqlx.Connect(
		"postgres",
		"host=localhost port=5432 user=admin password=passwd123 dbname=coderhouse sslmode=disable")
	if err != nil {
		panic(err)
	}

	r.Post("/users", createUser(db))

	http.ListenAndServe(":8080", r)
}
