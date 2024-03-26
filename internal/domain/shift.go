package domain


type Shift struct {
	ID           int       `json:"id" binding:"required"`
	PatientID     int       `json:"patient_id" binding:"required"`
	DentistID     int       `json:"dentist_id" binding:"required"`
	DischargeDate string    `json:"dischargedate" binding:"required"`
	Description   string    `json:"description" binding:"required"`
}
