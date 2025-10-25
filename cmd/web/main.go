package main

import (
	"log"
	"net/http"

	"github.com/obsidian33/contacts.app/internal/models"
)

type application struct {
	contacts models.ContactModelInterface
}

func main() {
	app := &application{
		contacts: &models.ContactModel{},
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: app.routes(),
	}

	log.Print("Starting server on :8080")
	err := srv.ListenAndServe()
	log.Fatal(err)
}
