package utils

import (
	"math/rand"
	"time"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	base     = len(alphabet)
)

func GenerateShortURL(num uint) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, num)
	for i := range b {
		b[i] = alphabet[rand.Intn(base)]
	}
	return string(b)
}
