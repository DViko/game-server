package helpers

import (
	"net/http"
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
		return nil
	}

	claims, ok := token.Claims.(*Claims)

	if !ok || !token.Valid {
		return nil
	}

	if claims.ExpiresAt.Time.Before(time.Now()) {
		return nil
	}

	return claims
}

func SetCookie(data RegisterUserResponse, w http.ResponseWriter) {

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    data.Token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Now().Add(time.Hour * 24),
	})
}

func ResetCookie(w http.ResponseWriter, r *http.Request) {

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})
}
