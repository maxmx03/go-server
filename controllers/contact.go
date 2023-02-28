package controllers

import (
	"fmt"
	"html/template"
	"net/http"
)

type Contact struct {
	Name    string
	Email   string
	Message string
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Render the contact from
		tmpl, err := template.ParseFiles("views/templates/contact.html")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, nil)
	} else if r.Method == "POST" {
		// handle Form submission
		name := r.FormValue("name")
		email := r.FormValue("email")
		message := r.FormValue("message")

		// Validate form fields
		if name == "" || email == "" || message == "" {
			fmt.Fprintf(w, "Please fill in all fields")
			return
		}

		// Send email

		// Render success message
		tmpl, err := template.ParseFiles("views/templates/contact-success.html")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		contact := Contact{Name: name, Email: email, Message: message}
		tmpl.Execute(w, contact)
	}
}
