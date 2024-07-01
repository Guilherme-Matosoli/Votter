package services

import (
	"database/sql"
	"fmt"

	"github.com/Guilherme-Matosoli/votter/internal/database/entity"
)

func GetPoll(db *sql.DB, poll_id string) (string, error) {
	var questionsList []*entity.Question

	questions, err := db.Query(`SELECT * FROM questions WHERE poll_id = $1`, poll_id)
	if err != nil {
		fmt.Println("Error happens in get_poll_service: ", err)
		return "Error", err
	}

	for questions.Next() {
		var question entity.Question

		err := questions.Scan(&question.Id, &question.Title, &question.Description)
		if err != nil {
			fmt.Println("Error happen in get_poll_service: ", err)
			return "Error", err
		}

		questionsList = append(questionsList, &question)
	}

	return "", nil
}
