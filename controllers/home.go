package controllers

import (
	"net/http"
	"example.com/go-server/views"
)

type homeData struct {
	Name string
	Age  int
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Create a new homeData object with the data to pass to the template
	data := homeData{Name: "Max", Age: 29}

	// Render the "home" template using the "views" package
	views.RenderTemplate(w, "home.html", data)
}
