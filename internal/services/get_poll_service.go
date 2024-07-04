package services

import (
	"database/sql"
	"fmt"
)

type poll struct {
	Id          string      `json:"id"`
	Title       string      `json:"title"`
	Total_Votes int         `json:"total_votes"`
	Questions   []*question `json:"questions"`
}

func GetPoll(db *sql.DB, poll_id string) (*poll, error) {
	poll := &poll{}

	questions, err := GetQuestion(db, poll_id)
	if err != nil {
		return poll, err
	}

	poll.Questions = questions

	for _, value := range questions {
		poll.Total_Votes += value.Votes
	}

	error := db.QueryRow(`SELECT id, title FROM polls WHERE id = $1`, poll_id).Scan(&poll.Id, &poll.Title)
	if error != nil {
		fmt.Println("Error happens on get_poll_service: ", err)
		return poll, error
	}

	return poll, nil
}
