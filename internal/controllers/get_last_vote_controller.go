package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Guilherme-Matosoli/votter/internal/database"
	"github.com/Guilherme-Matosoli/votter/internal/services"
	"github.com/Guilherme-Matosoli/votter/internal/utils"
)

func GetLastVote(w http.ResponseWriter, r *http.Request) {
	ip := utils.GetIp(r)

	conn, err := database.Connection()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&response{Message: "Internal Server Error"})
		return
	}
	defer conn.Close()

	success, err := services.GetLastVote(conn, ip)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&response{Message: "Internal Server Error"})
		return
	}

	if success {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(&response{Message: "Did not vote in the last 24h"})
		return
	}

	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(&response{Message: "Already vote in the last 24h"})
}
