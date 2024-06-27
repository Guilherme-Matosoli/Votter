package entity

import (
	"time"

	"github.com/Guilherme-Matosoli/votter/internal/utils"
	"github.com/google/uuid"
)

type Poll struct {
	Id         string    `json:"id"`
	Title      string    `json:"title"`
	Created_at time.Time `json:"created_at"`
}

func NewPoll(title string) *Poll {
	return &Poll{
		Id:         uuid.New().String(),
		Title:      title,
		Created_at: utils.GetTime(),
	}
}
