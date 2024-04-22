package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret")

func CreateToken(userId string) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func ParseToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return secretKey, nil
	})
	if err != nil {
		return "", fmt.Errorf("token parsing failed: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if userId, ok := claims["user_id"].(string); ok {
			return userId, nil
		}
		return "", fmt.Errorf("user_id not found in claims")
	}
	return "", fmt.Errorf("invalid token")
}
