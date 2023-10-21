package util

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const tokenTTL = 12 * time.Hour

type userTokenClaims struct {
	jwt.StandardClaims
	Id int `json:"id"`
	Role string `json:"role"`
}

func GenerateToken(id int, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &userTokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt: time.Now().Unix(),
		},
		id,
		role,
	})
	signingKey := os.Getenv("SIGNING_KEY")

	return token.SignedString([]byte(signingKey))
}

func ParseToken(accessToken string) (int, string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &userTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		signingKey := os.Getenv("SIGNING_KEY")

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, "", err
	}

	claims, ok := token.Claims.(*userTokenClaims)
	if !ok {
		return 0, "", errors.New("token claims are not of type *userTokenClaims")
	}

	return claims.Id,claims.Role, nil
}

func  GetTokenFromHeader(header string) (string, error) {
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
