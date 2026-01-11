package main

import (
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../../../ui/static"))))

	mux.HandleFunc("GET /", app.index)
	mux.HandleFunc("GET /contacts", app.contactList)
	//mux.HandleFunc("GET /contacts/new", app.contactCreate)
	//mux.HandleFunc("POST /contacts/new", app.contactCreatePost)

	return mux
}
