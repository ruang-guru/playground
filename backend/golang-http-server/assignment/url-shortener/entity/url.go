package entity

import (
	"crypto/sha256"
	"fmt"
	"time"
)

// Buatlah entity URL dengan atribut LongURL dan ShortURL masing-masing string

// TODO: answer here

func GetRandomShortURL(longURL string) string {
	s := fmt.Sprintf("%s%d", longURL, time.Now().Unix())
	sum := sha256.Sum256([]byte(s))
	hashString := fmt.Sprintf("%x", sum)
	return hashString[:8]
}
