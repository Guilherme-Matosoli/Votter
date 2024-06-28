package services

import (
	"database/sql"
	"fmt"

	"github.com/Guilherme-Matosoli/votter/internal/database/entity"
)

func CreateVote(db *sql.DB, props *entity.Vote) error {
	newVote := entity.NewVote(props.Ip_address, props.Poll_id)

	_, err := db.Exec("INSERT INTO votes (id, ip_address, voted_at, poll_id) VALUES ($1, $2, $3, $4)",
		newVote.Id, newVote.Ip_address, newVote.Voted_at, newVote.Poll_id)

	if err != nil {
		fmt.Println("Error has happen on create_vote_service: ", err)
	}

	return err
}
