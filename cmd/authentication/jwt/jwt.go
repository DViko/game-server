package jwt

import (
	"authentication/helpers"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var key = []byte("the_game_aythentication")

func GenerateToken(username string) string {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
		"iat":      time.Now().Unix(),
	})

	tokenString, err := token.SignedString(key)

	helpers.ErrorHelper(err, "Error creating token:")

	return tokenString
}
