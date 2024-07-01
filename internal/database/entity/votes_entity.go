package entity

import (
	"time"

	"github.com/Guilherme-Matosoli/votter/internal/utils"
	"github.com/google/uuid"
)

type Vote struct {
	Id         string    `json:"id"`
	Ip_address string    `json:"ip_address"`
	Voted_at   time.Time `json:"voted_at"`
	Voted_in   string    `json:"voted_in"`
	Poll_id    string    `json:"poll_id"`
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
