package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Guilherme-Matosoli/votter/internal/database"
	"github.com/Guilherme-Matosoli/votter/internal/database/entity"
	"github.com/Guilherme-Matosoli/votter/internal/services"
)

type createPollRequestBody struct {
	Title     string            `json:"title"`
	Questions []entity.Question `json:"questions"`
}

type response struct {
	Message string `json:"message"`
}

type pollid struct {
	Id string `json:"id"`
}

func CreatePollController(w http.ResponseWriter, r *http.Request) {
	conn, error := database.Connection()
	if error != nil {
		fmt.Println("Error try connect db: ", error)

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&response{Message: "Internal Server Error"})
		return
	}
	defer conn.Close()

	var input createPollRequestBody

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		fmt.Println("Error in create poll controller: ", err)

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&response{Message: "Internal Server Error"})
		return
	}

	if len(input.Questions) < 2 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(&response{Message: "Minimum of 2 questions"})
		return
	}

	id, err := services.CreatePoll(conn, input.Title, input.Questions)
	if err != nil {
		fmt.Println("Error in create poll controller: ", err)

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&response{Message: "Internal Server Error"})
		return
	}

	responseMessage, err := json.Marshal(&pollid{Id: id})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(responseMessage)
}
