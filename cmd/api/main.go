package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Guilherme-Matosoli/votter/internal/controllers"
	"github.com/Guilherme-Matosoli/votter/internal/database"
	"github.com/go-chi/chi"
)

type response struct {
	Ip string `json:"ip"`
}

func main() {
	database.RunMigration()
	fmt.Println("Database connection successfully")

	r := chi.NewRouter()

	r.Get("/ip", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Header.Get("x-real-ip"))
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&response{Ip: r.Header.Get("x-real-ip")})
	})

	r.Get("/poll/{id}", controllers.GetPoll)
	r.Get("/getvote/{id}", controllers.GetUserVote)
	r.Post("/poll/create", controllers.CreatePollController)
	r.Post("/vote/create", controllers.CreateVoteController)

	http.ListenAndServe(":4000", r)
	fmt.Println("Server is running on port 4000 :rocket:")
}
