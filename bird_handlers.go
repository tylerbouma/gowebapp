package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Bird struct {
	Species     string `json:"species"`
	Description string `json:"description"`
}

var birds []Bird

func getBirdHandler(w http.ResponseWriter, r *http.Request) {
	// Convert the "birds" variable to json
	birdListBytes, err := json.Marshal(birds)

	if err != nil {
		fmt.Println(fmt.Errorf("error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Write the json list of birds to the response
	w.Write(birdListBytes)
}

func createBirdHandler(w http.ResponseWriter, r *http.Request) {
	// New instance of bird
	bird := Bird{}

	// All data is in HTML for data. ParseForm parses the form values
	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Get the bird information
	bird.Species = r.Form.Get("species")
	bird.Description = r.Form.Get("description")

	birds = append(birds, bird)

	// Redirect the user to the original HTML page
	http.Redirect(w, r, "/assets/", http.StatusFound)
}
