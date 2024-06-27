package utils

import (
	"fmt"
	"time"
)

func GetTime() time.Time {
	location, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		fmt.Println("Error in get Brazil time: ", err)
	}

	timeNow := time.Now().In(location)

	return timeNow
}
