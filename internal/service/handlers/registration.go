package handlers

import (
	"html/template"
	"net/http"
)

func (h *Handler) Registration(login, pass string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/registration.gohtml"))

		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		login = r.FormValue("email")
		pass = r.FormValue("message")

		success := struct{ Success bool }{true}
		tmpl.Execute(w, success)
	}
}
