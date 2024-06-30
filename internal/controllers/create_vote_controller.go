package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Guilherme-Matosoli/votter/internal/database/entity"
	"github.com/Guilherme-Matosoli/votter/internal/utils"
)

type createVoteRequestBody struct {
	Ip_address string `json:"ip_address"`
	Poll_id    string `json:"poll_id"`
	Voted_in   string `json:"voted_in"`
}

func CreateVoteController(w http.ResponseWriter, r *http.Request) {
	ip := utils.GetIp(r)
	input := entity.Vote{Ip_address: ip}
	json.NewDecoder(r.Body).Decode(&input)

	fmt.Println(input)
}
