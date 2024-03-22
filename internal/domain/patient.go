package domain

import "time"

type Patient struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	LastName      string    `json:"lastname"`
	Home          string    `json:"home"`
	DNI           string    `json:"DNI"`
	DischargeDate time.Time `json:"dischargedate"`
}
