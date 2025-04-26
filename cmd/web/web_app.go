package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"time"
)

type RegisterUserResponse struct {
	UserID   string `json:"userId"`
	Username string `json:"username"`
	Token    string `json:"token"`
	Error    int32  `json:"error"`
}

func main() {

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("public/index.html"))
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/registration", registrationHandler)

	http.ListenAndServe(":8081", nil)
	log.Printf("Server started on port 8081")

}

func registrationHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		tmpl := template.Must(template.ParseFiles("public/registration.html"))
		tmpl.Execute(w, nil)
	}

	payload := map[string]string{
		"email":    r.FormValue("email"),
		"username": r.FormValue("username"),
		"password": r.FormValue("password"),
	}

	jData, err := json.Marshal(payload)
	if err != nil {
		log.Fatal("Json error", err)
	}

	resp, err := http.Post("http://localhost:8080/v1/authentication/registration", "application/json", bytes.NewBuffer(jData))
	if err != nil {
		log.Fatal("Response error", err)
	}

	defer resp.Body.Close()

	var result RegisterUserResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatal("Json error", err)
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    result.Token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Now().Add(time.Hour * 24),
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
