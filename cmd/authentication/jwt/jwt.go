package jwt

import (
	"authentication/helpers"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var key = []byte("the_game_aythentication")

func GenerateToken(uId string, uName string) string {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":   uId,
		"username": uName,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
		"iat":      time.Now().Unix(),
	})

	tokenString, err := token.SignedString(key)

	helpers.ErrorHelper(err, "Error creating token:")

	return tokenString
}
