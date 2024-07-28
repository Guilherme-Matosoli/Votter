package entity

import (
	"time"

	"github.com/Guilherme-Matosoli/votter/internal/utils"
)

type Poll struct {
	Id         string    `json:"id"`
	Title      string    `json:"title"`
	Created_at time.Time `json:"created_at"`
}

func NewPoll(title string) *Poll {
	return &Poll{
		Id:         utils.GenId(),
		Title:      title,
		Created_at: utils.GetTime(),
	}
}
