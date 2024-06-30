package entity

import (
	"time"

	"github.com/Guilherme-Matosoli/votter/internal/utils"
	"github.com/google/uuid"
)

type Vote struct {
	Id         string
	Ip_address string
	Voted_at   time.Time
	Voted_in   string
	Poll_id    string
}

func NewVote(ip_address string, poll_id string, voted_in string) *Vote {
	return &Vote{
		Id:         uuid.New().String(),
		Ip_address: ip_address,
		Voted_at:   utils.GetTime(),
		Voted_in:   voted_in,
		Poll_id:    poll_id,
	}
}
