package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"

	"web_app/helpers"
)

const sCrt = "certificate/server.crt"
const sKey = "certificate/server.key"

func main() {

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/signup", registrationHandler)

	err := http.ListenAndServeTLS(":8081", sCrt, sKey, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	log.Printf("Server started on port 8081")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("public/index.html"))
	tmpl.Execute(w, nil)
}

func registrationHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		payload := map[string]string{
			"email":    r.FormValue("email"),
			"username": r.FormValue("username"),
			"password": r.FormValue("password"),
		}

		jData := helpers.JsonEncoder(payload)

		resp, err := http.Post("https://localhost:8080/v1/authentication/signup", "application/json", bytes.NewBuffer(jData))
		if err != nil {
			log.Fatal("Response error", err)
		}

		helpers.SetCookie(helpers.JsonDecoder(resp), w)

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	tmpl := template.Must(template.ParseFiles("public/signup.html"))
	tmpl.Execute(w, nil)
}
