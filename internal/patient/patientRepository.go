package patient

import (
	"errors"
	"repositoryapi/internal/domain"
	"repositoryapi/pkg/store"
)

type PatientRepositoryInterface interface {
	ReadAllPatient() ([]domain.Patient, error)
	FindPatientById(id int) (domain.Patient, error)
	CreatePatient(patient domain.Patient) (domain.Patient, error)
	UpdatePatient(patient domain.Patient) (domain.Patient, error)
	PatchPatient(patient domain.Patient) (domain.Patient, error)
	DeletePatient(id int) error
}

type patientRepository struct {
	storage store.PatientStoreInterface
}

func NewPatientRepository(storage store.PatientStoreInterface) PatientRepositoryInterface {
	return &patientRepository{storage}
}

func (p *patientRepository) ReadAllPatient() ([]domain.Patient, error) {
	return p.storage.ReadAllPatient()
}

func (p *patientRepository) FindPatientById(id int) (domain.Patient, error) {
	patient, err := p.storage.ReadPatient(id)
	if err != nil {
		return domain.Patient{}, errors.New("patient not found")
	}
	return patient, nil
}

func (p *patientRepository) CreatePatient(patient domain.Patient) (domain.Patient, error) {
	return p.storage.CreatePatient(patient)
}

func (p *patientRepository) UpdatePatient(patient domain.Patient) (domain.Patient, error) {
	return p.storage.UpdatePatient(patient)
}

func (p *patientRepository) PatchPatient(patient domain.Patient) (domain.Patient, error) {
	return p.storage.PatchPatient(patient)
}

func (p *patientRepository) DeletePatient(id int) error {
	if !p.storage.ExistPatient(id) {
		return errors.New("Dentist not found")
	}
	return p.storage.DeletePatient(id)
}
