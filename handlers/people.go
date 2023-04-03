package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
)

func PeopleHandler(w http.ResponseWriter, r *http.Request) {
	// Simulate a database of people
	people := []models.Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 40},
		{Name: "Charlie", Age: 50},
	}

	// Convert people to JSON
	jsonBytes, err := json.Marshal(people)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Set the response content type to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	fmt.Fprint(w, string(jsonBytes))
}
