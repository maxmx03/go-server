package views

import (
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, templateName string, data interface{}) {
	tmpl, err := template.ParseFiles("views/templates/" + templateName)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
