package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"notable/datastore"
	"notable/models"
)

// GetAppointments get appointments for a doctor based on the day
// the URL query parameter is date=2023-04-05
func GetAppointments(ds *models.DataStore, w http.ResponseWriter, r *http.Request) {
	doctorID, _ := strconv.Atoi(r.URL.Query().Get("doctor_id"))
	date := r.URL.Query().Get("date")

	appointments := datastore.GetAppointments(doctorID, date, ds)
	json.NewEncoder(w).Encode(appointments)

	http.NotFound(w, r)
}

func DeleteAppointment(ds *models.DataStore, w http.ResponseWriter, r *http.Request) {
	doctorID, _ := strconv.Atoi(r.URL.Query().Get("doctor_id"))
	appointmentID, _ := strconv.Atoi(r.URL.Query().Get("appointment_id"))

	if err := datastore.DeleteAppointment(doctorID, appointmentID, ds); err != nil {
		http.NotFound(w, r)
		return
	}
}

func AddAppointment(ds *models.DataStore, w http.ResponseWriter, r *http.Request) {
	doctorID, _ := strconv.Atoi(r.URL.Query().Get("doctor_id"))

	var appointment models.Appointment
	json.NewDecoder(r.Body).Decode(&appointment)

	if err := datastore.AddAppointment(doctorID, appointment, ds); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}
