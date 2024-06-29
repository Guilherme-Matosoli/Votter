package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Guilherme-Matosoli/votter/internal/database"
	"github.com/Guilherme-Matosoli/votter/internal/database/entity"
	//"github.com/Guilherme-Matosoli/votter/internal/services"
)

type requestBody struct {
	Info      entity.Poll       `json:"info"`
	Questions []entity.Question `json:"questions"`
}

func CreatePollController(w http.ResponseWriter, r *http.Request) {
	_, error := database.Connection()
	if error != nil {
		fmt.Println("error happen: ", error)
		panic(error)
	}

	var input requestBody

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		fmt.Println("error has happen: ", err)
	}

	fmt.Println("infos: ", input.Info)
	fmt.Println("questions: ", input.Questions)

}
