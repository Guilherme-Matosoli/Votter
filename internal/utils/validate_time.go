package utils

import (
	"time"
)

func ValidateTime(timeNow time.Time, lastTime time.Time) bool {
	timeSubtraction := lastTime.Sub(timeNow)
	maxTime := 24 * time.Hour

	if timeSubtraction < maxTime {
		return false
	}

	return true
}
