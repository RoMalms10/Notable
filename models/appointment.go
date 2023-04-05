package models

type Appointment struct {
	UUID      string  `json:"id"`
	StartTime string  `json:"start_time"`
	EndTime   string  `json:"end_time"`
	Patient   Patient `json:"patient"`
	Kind      string  `json:"kind"` // New Patient or Follow up
}
