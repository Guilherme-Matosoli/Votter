package services

import (
	"database/sql"
	"fmt"

	"github.com/Guilherme-Matosoli/votter/internal/database/entity"
)

func CreateQuestion(db *sql.DB, props *entity.Question) (string, error) {
	newQuestion := entity.NewQuestion(props.Poll_id, props.Option)

	_, err := db.Exec("INSERT INTO questions (id, poll_id, description, title) VALUES ($1,$2,$3, $4)",
		newQuestion.Id, newQuestion.Poll_id, newQuestion.Option,
	)

	if err != nil {
		fmt.Println("Error happens on create_question_service: ", err)
	}

	return "Question created", err
}
