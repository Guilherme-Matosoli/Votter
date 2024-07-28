package utils

import (
	"net/http"
	"time"
)

func AntiRender() {
	channel := time.NewTicker(15 * time.Second)

	for {
		select {
		case <-channel.C:
			http.NewRequest("GET", "https://votter.onrender.com", nil)
		}
	}
}
