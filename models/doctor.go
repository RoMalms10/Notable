package models

type Doctor struct {
	UUID      string        `json:"id"`
	FirstName string        `json:"first_name"`
	LastName  string        `json:"last_name"`
	Calendar  []Appointment `json:"calendar"`
}
