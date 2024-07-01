package utils

import (
	"time"
)

func ValidateTime(lastTime time.Time) bool {
	timeSubtraction := lastTime.Sub(GetTime())
	maxTime := 24 * time.Hour

	if lastTime.Year() == 1 {
		return true
	}

	if timeSubtraction < maxTime {
		return false
	}

	return true
}
