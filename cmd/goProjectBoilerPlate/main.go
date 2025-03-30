package main

import (
	"fmt"
	"github.com/simorina/goProjectBoilerPlate/internal/routes"
	"net/http"
)

func main() {
	// Create a new router using the NewRouter function from the routes package
	router := routes.NewRouter()

	// Start the HTTP server on port 8080 and use the router to handle requests
	fmt.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
