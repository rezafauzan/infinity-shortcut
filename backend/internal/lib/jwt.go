package lib

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	UserId int `json:"user_id"`
	CartId int `json:"cart_id"`
	jwt.RegisteredClaims
}

func GenerateToken(userId int, cartId int) (string, error) {
	claims := CustomClaims{
		UserId: userId,
		CartId: cartId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("APP_SECRET")))

	if err != nil {
		return "", errors.New("Failed to login! : " + err.Error())
	}

	return tokenString, nil
}