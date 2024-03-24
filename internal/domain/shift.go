package domain

import "time"

type Shift struct {
	ID            int       `json:"id"`
	PatientID     int       `json:"patient_id"`
	DentistID     int       `json:"dentist_id"`
	DischargeDate time.Time `json:"dischargedate"`
	Description   string    `json:"description"`
}
