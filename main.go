package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")

	// Declare the static file directory
	staticFileDirectory := http.Dir("./assets/")

	// Declare the handler that routes requests to their respective filename
	// Remove the /assets prefix so we only look for files (ex. index.html)
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))

	// Matches all routes starting with "/assets/"
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
	r.HandleFunc("/bird", getBirdHandler).Methods("GET")
	r.HandleFunc("/bird", createBirdHandler).Methods("POST")
	return r
}

func main() {
	// Declare a new router
	r := newRouter()
	http.ListenAndServe(":8000", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
