package main

import (
	"net/http"
	"example.com/go-server/controllers"
)

func main() {
	// Register HTTP route handlers
	http.HandleFunc("/", controllers.HomeHandler)
	http.HandleFunc("/contact", controllers.ContactHandler)

	// Serve static files from the "static" directory
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

  // Start the HTTP server
  http.ListenAndServe(":8080", nil)
}
