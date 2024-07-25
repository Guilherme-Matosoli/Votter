package entity

import "github.com/google/uuid"

type Question struct {
	Id      string `json:"id"`
	Poll_id string `json:"poll_id"`
	Option  string `json:"option"`
}

func NewQuestion(poll_id string, option string) *Question {
	return &Question{
		Id:      uuid.New().String(),
		Poll_id: poll_id,
		Option:  option,
	}
}
