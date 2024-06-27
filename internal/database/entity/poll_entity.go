package entity

import (
	"time"

	"github.com/Guilherme-Matosoli/votter/internal/utils"
	"github.com/google/uuid"
)

type poll struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Created_at  time.Time `json:"created_at"`
	Total_votes int       `json:"total_votes"`
}

func NewPoll(title string) *poll {
	return &poll{
		Id:         uuid.New().String(),
		Title:      title,
		Created_at: utils.GetTime(),
	}
}
