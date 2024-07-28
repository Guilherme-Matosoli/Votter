package services

import (
	"database/sql"
	"fmt"

	"github.com/Guilherme-Matosoli/votter/internal/database/entity"
)

func CreateVote(db *sql.DB, props *entity.Vote) (*entity.Vote, error) {
	vote, err := GetUserVote(db, props.Ip_address, props.Poll_id)
	if err != nil {
		return nil, err
	}

	if vote != nil {
		return nil, nil
	}

	newVote := entity.NewVote(props.Ip_address, props.Poll_id, props.Voted_in)

	_, error := db.Exec("INSERT INTO votes (id, ip_address, voted_at, voted_in,poll_id) VALUES ($1, $2, $3, $4, $5)",
		newVote.Id, newVote.Ip_address, newVote.Voted_at, newVote.Voted_in, newVote.Poll_id)

	if error != nil {
		fmt.Println("Error in create_vote_service: ", err)
	}

	return newVote, error
}
