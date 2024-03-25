package domain

type Patient struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	LastName      string    `json:"lastname"`
	Home          string    `json:"home"`
	DNI           string    `json:"DNI"`
	DischargeDate string    `json:"dischargedate"`
}
