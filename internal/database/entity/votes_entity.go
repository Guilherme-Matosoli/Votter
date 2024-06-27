package entity

import "github.com/google/uuid"

type vote struct {
	Id         string
	Ip_address string
	Voted_at   string
	Poll_id    string
}

func NewVote(ip_address string, poll_id string) *vote {
	return &vote{
		Id:         uuid.New().String(),
		Ip_address: ip_address,
		Voted_at:   "",
		Poll_id:    poll_id,
	}
}
