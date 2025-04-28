package main

import (
	"net/http"
	"web_app/helpers"
	"web_app/routes"
)

const (
	sCrt = "certificate/server.crt"
	sKey = "certificate/server.key"
)

func main() {

	fs := http.FileServer(http.Dir("./public/css"))
	http.Handle("/public/css/", http.StripPrefix("/public/css/", fs))

	rHandler := routes.NewRoutesHandler(helpers.ReadConfig("config/config.yaml"))
	rHandler.PreparingRoutes()

	http.ListenAndServeTLS(":8081", sCrt, sKey, nil)
}
