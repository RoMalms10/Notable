package handlers

import (
	"encoding/json"
	"net/http"

	"notable/models"

	"notable/datastore"
)

func GetDoctors(ds *models.DataStore, w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(ds.Doctors)
}

// AddDoctor
func AddDoctor(ds *models.DataStore, w http.ResponseWriter, r *http.Request) {
	var doctor models.Doctor
	err := json.NewDecoder(r.Body).Decode(&doctor)
	if err != nil {
		http.Error(w, "couldn't parse doctor info", http.StatusBadRequest)
		return
	}

	doctor, err = datastore.AddDoctor(doctor, ds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(doctor)
}

func GetDoctor(ds *models.DataStore, w http.ResponseWriter, r *http.Request) {
	doctorID := r.URL.Query().Get("doctor_id")

	doctor, err := datastore.GetDoctor(doctorID, ds)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(doctor)
}
