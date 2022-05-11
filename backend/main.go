package main

import (
	"log"
	"net/http"
	"os"

	"github.com/brianloveswords/airtable"
	"github.com/ej-limited/auditions/handlers"
	"github.com/ej-limited/auditions/pkg/mail"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var (
	mailuser = os.Getenv("MAIL_USER")
	mailpass = os.Getenv("MAIL_PASS")
	apiKey   = os.Getenv("API_KEY")
)

func main() {

	r := mux.NewRouter()
	mc := mail.NewMailClient(mailuser, mailpass)
	c := airtable.Client{
		APIKey: apiKey,
		BaseID: "appk0cwYJZsoXanmK",
	}
	aHandler := handlers.NewAuditionHandler(mc, &c)
	r.HandleFunc("/register", aHandler.SignUP).Methods("POST")
	handler := cors.Default().Handler(r)
	log.Println("starting server on 0.0.0.0:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))

}
