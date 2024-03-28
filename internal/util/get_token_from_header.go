package util

import (
	"errors"
	"strings"
)

func GetTokenFromHeader(header string) (string, error) {
	if header == "" {
		return "", errors.New("empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		return "", errors.New("invalid auth header")
	}

	token := headerParts[1]

	return token, nil
}
