package helpers

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var key = []byte("the_game_aythentication")

type Claims struct {
	UserId   string `json:"userId"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func ValidateToken(tokenString string) *Claims {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		log.Println("Error parsing token:", err)
		return nil
	}

	claims, ok := token.Claims.(*Claims)

	if !ok || !token.Valid {
		log.Println("Invalid token")
		return nil
	}

	if claims.ExpiresAt.Time.Before(time.Now()) {
		log.Println("Token expired")
		return nil
	}

	return claims
}
