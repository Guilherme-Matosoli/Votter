package services

import (
	"database/sql"
	"fmt"

	"github.com/Guilherme-Matosoli/votter/internal/database/entity"
)

type poll struct {
	*entity.Poll
	Total_Votes int         `json:"total_votes"`
	Questions   []*question `json:"questions"`
}

func GetPoll(db *sql.DB, poll_id string) (*poll, error) {
	poll := &poll{Poll: &entity.Poll{}}

	questions, err := GetQuestion(db, poll_id)
	if err != nil {
		return poll, err
	}

	vas := *questions[0]

	fmt.Println("questions: ", vas)

	fmt.Println("Pass from getQuestion")

	poll.Questions = questions

	for range questions {
		poll.Total_Votes += 1
	}

	row, err := db.Query(`SELECT id, title, created_at FROM polls WHERE id = $1`, poll_id)
	if err != nil {
		fmt.Println("Error happens on get_poll_service: ", err)
		return poll, err
	}

	row.Scan(&poll.Id, &poll.Title, &poll.Created_at)

	if err = row.Err(); err != nil {
		fmt.Println("Error on row: ", err)
		return poll, err
	}

	return poll, nil
}
