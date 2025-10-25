package main

import (
	"net/http"

	"github.com/obsidian33/contacts.app/internal/models"
	"github.com/obsidian33/contacts.app/ui/components"
)

func (app *application) index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/contacts", http.StatusFound)
}

func (app *application) contactList(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("q")
	contacts := []models.Contact{}
	if search != "" {
		contacts, _ = app.contacts.Search(search)
	} else {
		contacts, _ = app.contacts.All()
	}
	messages := []string{}
	components.Render(w, r, components.Index(search, contacts, messages))
}

func (app *application) contactCreate(w http.ResponseWriter, r *http.Request) {
	components.Render(w, r, components.NewContact(models.Contact{}))
}

func (app *application) contactCreatePost(w http.ResponseWriter, r *http.Request) {
	c := models.Contact{
		First:  r.FormValue("first"),
		Last:   r.FormValue("last"),
		Phone:  r.FormValue("phone"),
		Email:  r.FormValue("email"),
		Errors: make(map[string]string),
	}

	if err := app.contacts.Save(c); err == nil {
		http.Redirect(w, r, "/contacts", http.StatusCreated)
	}

	components.Render(w, r, components.NewContact(c))
}
