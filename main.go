// Entrypoint for API
package main

import (
	"bookfeel/book"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {
	// Get the "PORT" env variable
	port := "8080"

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := book.NewRouter() // create routes

	// These two lines are important if you're designing a front-end to utilise this API methods
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	// Launch server with CORS validations
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(allowedOrigins, allowedMethods)(router)))
}
