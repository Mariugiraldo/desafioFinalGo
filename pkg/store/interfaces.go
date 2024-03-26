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
	UpdateShift(domain.Shift)(domain.Shift, error)
	DeleteShift(id int) 
	PatchShift(domain.Shift)(domain.Shift, error)
	/* CreateShiftByDni(dischargeDate string, description string,  dni string, registration string )(domain.Shift,error) */
	/* GetShiftsByPatientDNI(patientDNI string)([]domain.Shift, error)
 */
}

