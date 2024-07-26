package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Guilherme-Matosoli/votter/internal/database"
	"github.com/Guilherme-Matosoli/votter/internal/services"
	"github.com/go-chi/chi"
)

func GetPoll(w http.ResponseWriter, r *http.Request) {
	poll_id := chi.URLParam(r, "id")
	if poll_id == "" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(&response{Message: "Missing ID"})
	}

	conn, err := database.Connection()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&response{Message: "Internal Server Error"})

		fmt.Println("Error happens in get_poll_controller: ", err)
		return
	}
	defer conn.Close()

	poll, error := services.GetPoll(conn, poll_id)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&response{Message: "Internal Server Error"})

		fmt.Println("Error happens in get_poll_controller: ", error)
		return
	}

	if poll == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(&response{Message: "Poll doesn't exist"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(poll)
}
