package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Guilherme-Matosoli/votter/internal/database"
	"github.com/Guilherme-Matosoli/votter/internal/database/entity"
	"github.com/Guilherme-Matosoli/votter/internal/services"
	"github.com/Guilherme-Matosoli/votter/internal/utils"
)

type createVoteRequestBody struct {
	Ip_address string `json:"ip_address"`
	Poll_id    string `json:"poll_id"`
	Voted_in   string `json:"voted_in"`
}

func CreateVoteController(w http.ResponseWriter, r *http.Request) {
	conn, err := database.Connection()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&response{Message: "Internal Server Error"})

		fmt.Println("Error happens in create_vote_controller: ", err)
		return
	}
	defer conn.Close()

	ip := utils.GetIp(r)
	input := entity.Vote{Ip_address: ip}
	json.NewDecoder(r.Body).Decode(&input)

	msg, err := services.CreateVote(conn, &input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&response{Message: "Internal Server Error"})

		fmt.Println("Error happens in create_vote_controller: ", err)
		return

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(&response{Message: msg})
}
