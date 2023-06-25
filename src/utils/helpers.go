package utils

import (
	"math/rand"
	"net/url"
	"time"

	"github.com/skip2/go-qrcode"
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

func ValidateURL(inputURL string) bool {
	_, err := url.ParseRequestURI(inputURL)
	if err != nil {
		return false
	}
	return true
}

func GenerateQRCode(url string) ([]byte, error) {
	qrCode, err := qrcode.Encode(url, qrcode.Medium, 256)
	if err != nil {
		return nil, err
	}

	return qrCode, nil
}
