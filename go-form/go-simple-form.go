package main

import (
	"net/http"
	"text/template"
)

func main() {
	emailForm()
}

func emailForm() {

	type ContactDetails struct {
		Email   string
		Subject string
		Message string
	}

	tmpl := template.Must(template.ParseFiles("forms.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		details := ContactDetails{
			Email:   r.FormValue("email"),
			Subject: r.FormValue("subject"),
			Message: r.FormValue("message"),
		}

		// do something with details
		_ = details

		//tmpl.Execute(w, struct{ Success bool }{true})
		tmpl.Execute(w, struct {
			Success bool
			Email   string
		}{
			Success: true,
			Email:   details.Email,
		})
	})

	http.ListenAndServe(":8080", nil)
}
