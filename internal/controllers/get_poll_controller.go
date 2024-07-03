package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Guilherme-Matosoli/votter/internal/database"
)

func GetPoll(w http.ResponseWriter, r *http.Request) {
	conn, err := database.Connection()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&response{Message: "Internal Server Error"})

		fmt.Println("Error happens in create_vote_controller: ", err)
		return
	}
	defer conn.Close()

}
