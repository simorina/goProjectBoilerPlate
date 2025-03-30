package routes

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome, to the homepage!")
}

func apidataHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the API data")
}

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	// Register the index handler with the default route
	mux.HandleFunc("/", indexHandler)

	// Register the API data handler with the "/api/data" route
	mux.HandleFunc("/api/data", apidataHandler)

	return mux
}
