package store

import "repositoryapi/internal/domain"

type StoreInterface interface {
	Read(id int) (domain.Dentist, error)
	CreateDentist(dentist domain.Dentist) (domain.Dentist, error)
	UpdateDentist(dentist domain.Dentist) (domain.Dentist, error)
	PatchDentist(dentist domain.Dentist) (domain.Dentist, error)
	Delete(id int)

	ReadShift(id int) (domain.Shift, error)
	CreateShift(domain.Shift) (domain.Shift, error)
	UpdateShift(domain.Shift) (domain.Shift, error)
	DeleteShift(id int)
	PatchShift(domain.Shift) (domain.Shift, error)
}

type PatientStoreInterface interface {
	ReadAllPatient() ([]domain.Patient, error)
	ReadPatient(id int) (domain.Patient, error)
	CreatePatient(patient domain.Patient) (domain.Patient, error)
	UpdatePatient(patient domain.Patient) (domain.Patient, error)
	PatchPatient(patient domain.Patient) (domain.Patient, error)
	DeletePatient(id int) error
	ExistPatient(id int) bool
}
