package helpers

import (
	"net/http"
	"time"
)

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
