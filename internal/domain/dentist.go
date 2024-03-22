package domain

type Dentist struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	LastName     string `json:"lastname"`
	Registration string `json:"registration"`
}
