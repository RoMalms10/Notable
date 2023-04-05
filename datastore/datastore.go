package datastore

import (
	"fmt"
	"time"

	"notable/models"
)

func DeleteAppointment(doctorID int, appointmentID int, ds *models.DataStore) error {
	for i, doctor := range ds.Doctors {
		if doctor.UUID == doctorID {
			for j, appointment := range doctor.Calendar {
				if appointment.UUID == appointmentID {
					ds.Doctors[i].Calendar = append(doctor.Calendar[:j], doctor.Calendar[j+1:]...)
					return nil
				}
			}
		}
	}

	return fmt.Errorf("Couldn't find appointment")
}

func AddAppointment(doctorID int, appointment models.Appointment, ds *models.DataStore) error {
	for i, doctor := range ds.Doctors {
		if doctor.UUID == doctorID {
			if len(doctor.Calendar) >= 3 {
				return fmt.Errorf("Maximum number of appointments reached for this doctor")

			}

			if !isValidAppointment(doctor.Calendar, appointment) {
				return fmt.Errorf("Invalid appointment time")
			}

			appointment.UUID = len(doctor.Calendar) + 1
			ds.Doctors[i].Calendar = append(doctor.Calendar, appointment)
			// json.NewEncoder(w).Encode(appointment)
			return nil
		}
	}

	return fmt.Errorf("error scheduling appointment")
}

func isValidAppointment(calendar []models.Appointment, appointment models.Appointment) bool {
	// check appointment start time is at 15-minute intervals
	t, err := time.Parse("2006-01-02 15:04:05", appointment.StartTime)

	// make sure the start time is in 15 minute format
	if err != nil || t.Minute()%15 != 0 || t.Second() != 0 {
		return false
	}

	// check there are no more than 3 appointments at the same time for a doctor
	count := 0
	for _, a := range calendar {
		if a.StartTime == appointment.StartTime {
			count++
		}
	}
	if count >= 3 {
		return false
	}

	return true
}

func GetAppointments(doctorID int, date string, ds *models.DataStore) []models.Appointment {
	for _, doctor := range ds.Doctors {
		if doctor.UUID == doctorID {
			var appointments []models.Appointment
			for _, appointment := range doctor.Calendar {
				// grab the appointment time string up to the date
				if appointment.StartTime[:10] == date {
					appointments = append(appointments, appointment)
				}
			}

			return appointments
		}
	}
	return nil
}
