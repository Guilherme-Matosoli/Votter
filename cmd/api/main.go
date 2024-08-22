package main

import (
	"fmt"
	"net/http"

	"github.com/Guilherme-Matosoli/votter/internal/controllers"
	"github.com/Guilherme-Matosoli/votter/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func main() {
	database.RunMigration()
	fmt.Println("Database connection successfully")

	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
	}))

	r.Get("/poll/{id}", controllers.GetPoll)
	r.Get("/getvote/{id}", controllers.GetUserVote)
	r.Post("/poll/create", controllers.CreatePollController)
	r.Post("/vote/create", controllers.CreateVoteController)

	http.ListenAndServe(":4000", r)
	fmt.Println("Server is running on port 4000 :rocket:")
}
