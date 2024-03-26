package domain

type Dentist struct {
	Id           int    `json:"id" binding:"required"`
	Name         string `json:"name" binding:"required"`
	LastName     string `json:"lastname" binding:"required"`
	Registration string `json:"registration" binding:"required"`
}
