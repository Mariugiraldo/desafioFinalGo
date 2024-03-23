package domain

import "time"

type Shift struct {
	ID           int       `json:"id" binding:"required"`
	PatientID     int       `json:"patient_id" binding:"required"`
	DentistID     int       `json:"dentist_id" binding:"required"`
	DischargeDate time.Time `json:"dischargedate" binding:"required"`
	Description   string    `json:"description" binding:"required"`
}
