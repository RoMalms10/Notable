package datastore

import (
	"fmt"
	"time"

	"notable/models"

	"github.com/google/uuid"
)

func DeleteAppointment(doctorID string, appointmentID string, ds *models.DataStore) error {
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

func AddAppointment(doctorID string, appointment models.Appointment, ds *models.DataStore) error {
	for i, doctor := range ds.Doctors {
		if doctor.UUID == doctorID {

			_, err := isValidAppointment(doctor.Calendar, appointment)
			if err != nil {
				return err
			}

			appointment.UUID = uuid.New().String()
			ds.Doctors[i].Calendar = append(doctor.Calendar, appointment)
			return nil
		}
	}

	return fmt.Errorf("error scheduling appointment")
}

func isValidAppointment(calendar []models.Appointment, appointment models.Appointment) (bool, error) {
	// check appointment start time is at 15-minute intervals
	t, err := time.Parse("2006-01-02 15:04:05", appointment.StartTime)

	// make sure the start time is in 15 minute format
	if err != nil || t.Minute()%15 != 0 || t.Second() != 0 {
		return false, fmt.Errorf("invalid appointment time")
	}

	// check there are no more than 3 appointments at the same time for a doctor
	count := 0
	for _, a := range calendar {
		if a.StartTime == appointment.StartTime {
			count++
		}
	}
	if count >= 3 {
		return false, fmt.Errorf("maximum appointments reached")
	}

	return true, nil
}

func GetAppointments(doctorID string, date string, ds *models.DataStore) []models.Appointment {
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

func GetDoctor(doctorID string, ds *models.DataStore) (models.Doctor, error) {
	for _, d := range ds.Doctors {
		if d.UUID == doctorID {
			return d, nil
		}
	}

	return models.Doctor{}, fmt.Errorf("doctor not found")
}

func AddDoctor(doctor models.Doctor, ds *models.DataStore) (models.Doctor, error) {
	doctor.UUID = uuid.New().String()

	// check if doctor ID already exists
	for _, d := range ds.Doctors {
		if d.UUID == doctor.UUID {
			return models.Doctor{}, fmt.Errorf("doctor ID already exists")
		}
	}

	// add new doctor to data store
	ds.Doctors = append(ds.Doctors, doctor)
	return doctor, nil
}
