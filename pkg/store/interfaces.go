package store

import "repositoryapi/internal/domain"

type StoreInterface interface {
	Read(id int) (domain.Dentist, error)
	CreateDentist(dentist domain.Dentist) (domain.Dentist, error)
	UpdateDentist(dentist domain.Dentist) (domain.Dentist, error)
	PatchDentist(dentist domain.Dentist) (domain.Dentist, error)
	Delete(id int)
}

type PatientStoreInterface interface {
	Read(id int) (domain.Patient, error)
	CreatePatient(patient domain.Patient) (domain.Patient, error)
	UpdatePatient(patient domain.Patient) (domain.Patient, error)
	PatchPatient(patient domain.Patient) (domain.Patient, error)
	Delete(id int)
}

type ShiftStoreInterface interface {
	Read(id int) (domain.Shift, error)
	CreateShift(shift domain.Shift) (domain.Shift, error)
	UpdateShift(shift domain.Shift) (domain.Shift, error)
	PatchShift(shift domain.Shift) (domain.Shift, error)
	Delete(id int)
}
