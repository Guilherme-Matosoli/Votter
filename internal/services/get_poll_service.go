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

func GetPoll(db *sql.DB, poll_id string) (string, error) {
	var poll *pollInfos

	questions, err := GetQuestion(db, poll_id)
	if err != nil {
		return "", err
	}

	poll.Questions = questions

	for _, _ = range questions {
		poll.Poll.Total_Votes += 1
	}

	row, err := db.Query(`SELECT * FROM polls WHERE id = $1`, poll_id)
	if err != nil {
		fmt.Println("Error happens on get_poll_service: ", err)
		return "Error", err
	}

	row.Scan(&poll)

	return "", nil
}
