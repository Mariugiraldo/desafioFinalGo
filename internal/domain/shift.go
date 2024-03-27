package domain

type Shift struct {
	ID            int    `json:"id" `
	PatientID     int    `json:"patient_id"`
	DentistID     int    `json:"dentist_id"`
	DischargeDate string `json:"dischargedate" binding:"required"`
	Description   string `json:"description" binding:"required"`
}
