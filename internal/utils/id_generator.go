package utils

import (
	"bytes"
	"math/rand"
)

func GenId() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	var id bytes.Buffer

	for range 15 {
		id.WriteString(string(charset[rand.Intn(len(charset))]))
	}

	return id.String()
}
