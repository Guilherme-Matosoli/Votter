package services

import (
	"database/sql"
	"fmt"

	"github.com/Guilherme-Matosoli/votter/internal/database/entity"
)

type question struct {
	*entity.Question
	Votes int `json:"votes"`
}

func GetQuestion(db *sql.DB, poll_id string) ([]*question, error) {
	var questionsList []*question

	query := `SELECT *, votes
	 					FROM questions
						LEFT JOIN votes
						ON votes.voted_in = questions.Id
						WHERE questions.poll_id = $1`
	questions, err := db.Query(query, poll_id)

	if err != nil {
		fmt.Println("Error happens in get_question_service: ", err)
		return nil, err
	}

	for questions.Next() {
	}

	return questionsList, nil
}
