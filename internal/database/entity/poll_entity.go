package entity

import "github.com/google/uuid"

type poll struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Created_at  string `json:"created_at"`
	Total_votes int    `json:"total_votes"`
}

func NewPoll(description string, title string, created_at string) *poll {
	return &poll{
		Id:         uuid.New().String(),
		Title:      title,
		Created_at: created_at,
	}
}
