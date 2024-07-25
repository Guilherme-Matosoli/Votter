package services

import (
	"database/sql"
	"fmt"

	"github.com/Guilherme-Matosoli/votter/internal/database/entity"
)

func GetUserVote(db *sql.DB, poll_id string, ip string) (*entity.Vote, error) {
	var lastVote *entity.Vote

	err := db.QueryRow(`SELECT * FROM votes WHERE ip_address = $1`, ip).Scan(&lastVote)
	if err != nil {
		fmt.Println("Error in get user vote: ", err)
		return nil, err
	}

	return lastVote, nil
}
