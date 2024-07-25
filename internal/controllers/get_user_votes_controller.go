package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Guilherme-Matosoli/votter/internal/database"
	"github.com/Guilherme-Matosoli/votter/internal/services"
	"github.com/Guilherme-Matosoli/votter/internal/utils"
	"github.com/go-chi/chi"
)

func GetUserVote(w http.ResponseWriter, r *http.Request) {
	ip := utils.GetIp(r)
	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(&response{Message: "Missing fields"})
	}

	conn, err := database.Connection()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&response{Message: "Internal Server Error"})
	}

	vote, err := services.GetUserVote(conn, id, ip)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&response{Message: "Internal Server Error"})
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vote)
}
