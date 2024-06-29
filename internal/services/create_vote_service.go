package services

import (
	"database/sql"
	"fmt"

	"github.com/Guilherme-Matosoli/votter/internal/database/entity"
)

func CreateVote(db *sql.DB, props *entity.Vote) error {
	var lastVote entity.Vote
	row, err := db.Query(`SELECT "voted_at" FROM votes WHERE "id" = $1`, props.Ip_address)
	if err != nil {
		fmt.Println("Error happens: ", err)
	}
	row.Scan(&lastVote)

	newVote := entity.NewVote(props.Ip_address, props.Poll_id)

	_, error := db.Exec("INSERT INTO votes (id, ip_address, voted_at, poll_id) VALUES ($1, $2, $3, $4)",
		newVote.Id, newVote.Ip_address, newVote.Voted_at, newVote.Poll_id)

	if error != nil {
		fmt.Println("Error has happen on create_vote_service: ", err)
	}

	return err
}
