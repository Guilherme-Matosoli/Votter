package services

import (
	"database/sql"
	"fmt"

	"github.com/Guilherme-Matosoli/votter/internal/database/entity"
)

type question struct {
	*entity.Question
	Votes []*entity.Vote `json:"votes"`
}

func GetQuestion(db *sql.DB, poll_id string) ([]*question, error) {
	var questionsList []*question

	questions, err := db.Query(`SELECT *, votes 
	 														FROM questions
															LEFT JOIN votes
															ON votes.voted_in = questions.Id
															WHERE poll_id = $1`, poll_id)

	if err != nil {
		fmt.Println("Error happens in get_poll_service: ", err)
		return nil, err
	}

	for questions.Next() {
		var question question

		fmt.Println(question.Votes)

		err := questions.Scan(&question.Id, &question.Title, &question.Description, &question.Votes)
		if err != nil {
			fmt.Println("Error happen in get_poll_service: ", err)
			return nil, err
		}

		questionsList = append(questionsList, &question)
	}

	return questionsList, nil
}
