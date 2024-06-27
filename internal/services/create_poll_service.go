package services

import (
	"database/sql"
	"fmt"

	"github.com/Guilherme-Matosoli/votter/internal/database/entity"
)

type pollPropertys struct {
	Title string `json:"title"`
}

func CreatePoll(db *sql.DB, poll pollPropertys) (string, error) {
	newPoll := entity.NewPoll(poll.Title)

	_, err := db.Exec("INSERT INTO polls (id, title, created_at) VALUES ($1,$2,$3)", newPoll.Id, newPoll.Title, newPoll.Created_at)
	if err != nil {
		fmt.Println("Some error happens: ", err)
	}

	return "Poll created", err
}
