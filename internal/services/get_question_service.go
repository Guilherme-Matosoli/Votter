package services

import (
	"database/sql"
	"fmt"
)

type question struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Title       string `json:"title"`
	Votes       int    `json:"votes"`
}

func GetQuestion(db *sql.DB, poll_id string) ([]*question, error) {
	var questionsList []*question

	query := `SELECT questions.id, questions.title, questions.description, COUNT(votes.id) as votes
	 					FROM questions
						LEFT JOIN votes
						ON votes.voted_in = questions.Id
						WHERE questions.poll_id = $1
						GROUP BY questions.id`
	questions, err := db.Query(query, poll_id)

	if err != nil {
		fmt.Println("Error happens in get_question_service: ", err)
		return nil, err
	}

	for questions.Next() {
		question := &question{}

		err := questions.Scan(&question.Id, &question.Title, &question.Description, &question.Votes)
		if err != nil {
			fmt.Println("Error happens in get_question_service (loop quesitons) ", err)
		}

		questionsList = append(questionsList, question)
	}

	if err = questions.Err(); err != nil {
		fmt.Println("erro happens: ", err)
		return nil, err
	}

	return questionsList, nil
}
