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

type vote struct {
	Votes []*entity.Vote
}

func GetQuestion(db *sql.DB, poll_id string) ([]*question, error) {
	var questionsList []*question

	query := `SELECT *, votes
	 					FROM questions
						LEFT JOIN votes
						ON votes.voted_in = questions.Id
						WHERE poll_id = $1`
	questions, err := db.Query(query, poll_id)

	if err != nil {
		fmt.Println("Error happens in get_poll_service: ", err)
		return nil, err
	}

	for questions.Next() {
		var question question
		var votes vote

		err := questions.Scan(&question.Id, &question.Title, &question.Description, &question.Votes, &votes.Votes)
		if err != nil {
			fmt.Println("Error happen in get_poll_service: ", err)
			return nil, err
		}

		question.Votes = len(votes.Votes)
		questionsList = append(questionsList, &question)
	}

	return questionsList, nil
}
