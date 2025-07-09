package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const userSecretKey = "hellouser12369420"

func GenerateToken(user_id int64, email, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"role":    role,
		"email":   email,
		"user_id": user_id,
		"exp":     time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(userSecretKey))
}

func VerifiedToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signin method")
		}

		return []byte(userSecretKey), nil
	})
	if err != nil {
		return 0, errors.New("couldn't parse token")
	}

	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		return 0, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	user_id := int64(claims["user_id"].(float64))
	return user_id, nil
}
