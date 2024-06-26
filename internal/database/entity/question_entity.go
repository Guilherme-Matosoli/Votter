package entity

import "github.com/google/uuid"

type question struct {
	Id          string `json:"id"`
	Poll_id     string `json:"poll_id"`
	Description string `json:"description"`
}

func NewQuestion(poll_id string, description string) *question {
	return &question{
		Id:          uuid.New().String(),
		Poll_id:     poll_id,
		Description: description,
	}
}
