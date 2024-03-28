package util

import (
	"crypto/sha1"
	"fmt"
	"os"
)

func GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	salt := os.Getenv("SALT")

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}