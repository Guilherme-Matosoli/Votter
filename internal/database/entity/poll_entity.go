package entity

import "github.com/google/uuid"

type poll struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Title       string `json:"title"`
	Created_at  string `json:"created_at`
	Total_votes int    `json:"total_votes"`
}

func CreatePoll(description string, title string, created_at string) *poll {
	return &poll{
		Id:          uuid.New().String(),
		Description: description,
		Title:       title,
		Created_at:  created_at,
	}
}
