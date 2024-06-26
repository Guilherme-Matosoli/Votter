package entity

import "github.com/google/uuid"

type poll struct {
	Id          string
	Description string
	Title       string
	Created_at  string
	Total_votes int
}

func CreatePoll(description string, title string, created_at string) *poll {
	return &poll{
		Id:          uuid.New().String(),
		Description: description,
		Title:       title,
		Created_at:  created_at,
	}
}
