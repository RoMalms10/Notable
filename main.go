package main

import (
	"net/http"

	"notable/handlers"
	"notable/models"
)

// if I end up using the other folders: import "../handlers"

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var people []Person

func main() {
	dataStore := NewDataStore()

	http.HandleFunc("/doctors", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetDoctors(dataStore, w, r)
	})

	http.HandleFunc("/doctor", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetDoctor(dataStore, w, r)
	})

	http.HandleFunc("/doctor/add", func(w http.ResponseWriter, r *http.Request) {
		handlers.AddDoctor(dataStore, w, r)
	})

	http.HandleFunc("/appointments", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			handlers.GetAppointments(dataStore, w, r)
		case "DELETE":
			handlers.DeleteAppointment(dataStore, w, r)
		case "POST":
			handlers.AddAppointment(dataStore, w, r)
		default:
			http.Error(w, "Unsupported HTTP method", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8080", nil)
}

func NewDataStore() *models.DataStore {
	return &models.DataStore{}
}
