package main

import (
	"net/http"

	"github.com/a-h/templ"
)

func (app *application) render(w http.ResponseWriter, r *http.Request, component templ.Component) error {
	return component.Render(r.Context(), w)
}
