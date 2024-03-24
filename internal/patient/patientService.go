package patient

import (
	"errors"
	"repositoryapi/internal/domain"
)

type PatientServiceInterface interface {
	ReadAllPatient() ([]domain.Patient, error)
	FindPatientById(id int) (domain.Patient, error)
	CreatePatient(patient domain.Patient) (domain.Patient, error)
	UpdatePatient(patient domain.Patient) (domain.Patient, error)
	PatchPatient(patient domain.Patient) (domain.Patient, error)
	DeletePatient(id int) error
}

type patientService struct {
	r PatientRepositoryInterface
}

func NewPatientService(r PatientRepositoryInterface) PatientServiceInterface {
	return &patientService{r}
}

func (service *patientService) ReadAllPatient() ([]domain.Patient, error) {
	patients, err := service.r.ReadAllPatient()
	if err != nil {
		return nil, err
	}
	return patients, nil
}

func (service *patientService) FindPatientById(id int) (domain.Patient, error) {
	patient, err := service.r.FindPatientById(id)
	if err != nil {
		return domain.Patient{}, err
	}
	return patient, nil
}

func (service *patientService) CreatePatient(patient domain.Patient) (domain.Patient, error) {
	patient, err := service.r.CreatePatient(patient)
	if err != nil {
		return domain.Patient{}, err
	}
	return patient, nil
}

func (service *patientService) UpdatePatient(pat domain.Patient) (domain.Patient, error) {
	if pat.Name == "" || pat.LastName == "" || pat.Home == "" || pat.DNI == "" {
		return domain.Patient{}, errors.New("all fields are required")
	}
	patient, err := service.r.UpdatePatient(pat)
	if err != nil {
		return domain.Patient{}, err
	}
	return patient, nil
}

func (service *patientService) PatchPatient(patient domain.Patient) (domain.Patient, error) {
	if patient.Name == "" {
		return domain.Patient{}, errors.New("name is required")
	}
	patient, err := service.r.PatchPatient(patient)
	if err != nil {
		return domain.Patient{}, err
	}
	return patient, nil
}

func (service *patientService) DeletePatient(id int) error {
	return service.r.DeletePatient(id)
}
