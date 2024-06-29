package main

import (
	"net/http"

	"github.com/Guilherme-Matosoli/votter/internal/database"
	"github.com/go-chi/chi"
)

func main() {
	database.RunMigration()

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world"))
	})

	http.ListenAndServe(":4000", r)
}
