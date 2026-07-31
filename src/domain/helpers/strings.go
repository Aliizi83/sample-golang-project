package helpers

import (
	"math/rand"
	"time"
)

const charSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_"

func GenerateRandomString(length int) string {
	if length <= 0 {
		return ""
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]byte, length)
	for i := range result {
		result[i] = charSet[r.Intn(len(charSet))]
	}

	return string(result)
}
