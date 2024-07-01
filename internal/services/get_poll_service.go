package services

import (
	"database/sql"
	"fmt"

	"github.com/Guilherme-Matosoli/votter/internal/database/entity"
)

type pollInfos struct {
	Poll      *entity.Poll `json:"poll"`
	Questions []*question  `json:"questions"`
}

func GetPoll(db *sql.DB, poll_id string) (string, error) {
	var poll *entity.Poll

	row, err := db.Query(`SELECT * FROM polls WHERE id = $1`, poll_id)
	if err != nil {
		fmt.Println("Error happens on get_poll_service: ", err)
		return "Error", err
	}

	row.Scan(&poll)

	return "", nil
}
