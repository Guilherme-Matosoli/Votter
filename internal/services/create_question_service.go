package services

import (
	"database/sql"
	"fmt"

	"github.com/Guilherme-Matosoli/votter/internal/database/entity"
)

type questionPropertys struct {
	Poll_id     string
	Description string
	Title       string
}

func CreateQuestion(db *sql.DB, props *questionPropertys) (string, error) {
	newQuestion := entity.NewQuestion(props.Poll_id, props.Description, props.Title)

	_, err := db.Exec("INSERT INTO questions (id, poll_id, title, description, title) VALUES ($1,$2,$3, $4)",
		newQuestion.Id, newQuestion.Poll_id, newQuestion.Description, newQuestion.Title,
	)

	if err != nil {
		fmt.Println("Error happens on create_question_service: ", err)
	}

	return "Question created", err
}
