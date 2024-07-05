package services

import (
	"database/sql"
	"fmt"

	"github.com/Guilherme-Matosoli/votter/internal/database/entity"
	"github.com/Guilherme-Matosoli/votter/internal/utils"
)

func GetLastVote(db *sql.DB, ip string) (bool, error) {
	var lastVote entity.Vote
	err := db.QueryRow(`SELECT * FROM votes WHERE "ip_address" = $1`, ip).Scan(&lastVote.Id, &lastVote.Ip_address, &lastVote.Voted_at, &lastVote.Voted_in, &lastVote.Poll_id)
	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Println("Erro happens on create_vote_service: ", err)
			return false, nil
		}

		return true, nil
	}

	validTimeToVote := utils.ValidateTime(lastVote.Voted_at)

	if !validTimeToVote {
		return false, err
	}

	return true, nil
}
