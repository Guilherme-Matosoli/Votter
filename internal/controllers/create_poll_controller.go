package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Guilherme-Matosoli/votter/internal/database"
	"github.com/Guilherme-Matosoli/votter/internal/database/entity"
	"github.com/Guilherme-Matosoli/votter/internal/services"
)

type requestBody struct {
	Info      entity.Poll       `json:"info"`
	Questions []entity.Question `json:"questions"`
}

type response struct {
	Message string `json:"message"`
}

func CreatePollController(w http.ResponseWriter, r *http.Request) {
	conn, error := database.Connection()
	if error != nil {
		fmt.Println("error happen: ", error)

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response{Message: "Internal Server Error"})
		return
	}

	var input requestBody

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response{Message: "Internal Server Error"})
		return
	}

	msg, err := services.CreatePoll(conn, &input.Info, input.Questions)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response{Message: "Internal Server Error"})
		return
	}

	responseMessage, err := json.Marshal(response{Message: msg})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(responseMessage)
}
