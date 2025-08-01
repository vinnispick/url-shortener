package service

import "crypto/rand"

func GenerateShortUrl() string {
	return rand.Text()
}
