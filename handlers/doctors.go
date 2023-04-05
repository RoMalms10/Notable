package handlers

import (
	"encoding/json"
	"net/http"

	"notable/models"

	"github.com/google/uuid"
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
	doctor.ID = uuid.New().String()

	// check if doctor ID already exists
	for _, d := range ds.doctors {
		if d.ID == doctor.ID {
			return // doctor ID already exists, do nothing
		}
	}
	// assign unique ID to new doctor
	if len(ds.doctors) > 0 {
		doctor.ID = ds.doctors[len(ds.doctors)-1].ID + 1
	} else {
		doctor.ID = 1
	}

	// add new doctor to data store
	ds.doctors = append(ds.doctors, doctor)
	json.NewEncoder(w).Encode(doctor)
}

func GetDoctor(ds *models.DataStore, w http.ResponseWriter, r *http.Request) {

}
