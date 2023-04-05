package models

type Patient struct {
	UUID      int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
