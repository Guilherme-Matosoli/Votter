package services

import (
	"database/sql"
	"fmt"

	"github.com/Guilherme-Matosoli/votter/internal/database/entity"
)

func CreatePoll(db *sql.DB, poll *entity.Poll, questions []*entity.Question) (string, error) {
	newPoll := entity.NewPoll(poll.Title)

	_, err := db.Exec("INSERT INTO polls (id, title, created_at) VALUES ($1,$2,$3)", newPoll.Id, newPoll.Title, newPoll.Created_at)
	if err != nil {
		fmt.Println("Some error happens: ", err)
	}

	for _, question := range questions {
		props := &entity.Question{Poll_id: question.Poll_id, Description: question.Description, Title: question.Title}

		_, err := CreateQuestion(db, props)
		if err != nil {
			fmt.Println("Error has happen on create_poll_service: ", err)
		}
	}

	return "Poll created", err
}
