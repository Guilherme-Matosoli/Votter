package services

import (
	"database/sql"
	"fmt"

	"github.com/Guilherme-Matosoli/votter/internal/database/entity"
)

type poll struct {
	*entity.Poll
	Total_Votes int `json:"total_votes"`
}

type pollInfos struct {
	Poll      poll        `json:"poll"`
	Questions []*question `json:"questions"`
}

func GetPoll(db *sql.DB, poll_id string) error {
	var poll *pollInfos

	questions, err := GetQuestion(db, poll_id)
	if err != nil {
		return err
	}

	fmt.Println("Pass from getQuestion")

	poll.Questions = questions

	for range questions {
		poll.Poll.Total_Votes += 1
	}

	row, err := db.Query(`SELECT * FROM polls WHERE id = $1`, poll_id)
	if err != nil {
		fmt.Println("Error happens on get_poll_service: ", err)
		return err
	}

	row.Scan(&poll)

	return nil
}
