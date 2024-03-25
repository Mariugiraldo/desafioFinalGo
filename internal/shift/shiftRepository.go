package shift

import (
	"errors"
	"repositoryapi/internal/domain"
	"repositoryapi/pkg/store"
)

type ShiftRepository interface {
	GetByID(id int) (domain.Shift, error)
	CreateShift(domain.Shift) (domain.Shift, error)
	UpdateShift(domain.Shift) (domain.Shift, error)
	DeleteShift(id int)
	PatchShift(shift domain.Shift) (domain.Shift, error)
	CreateShiftByDni(dischargeDate string, description string, dni string, registration string) (domain.Shift, error)
	/* GetShiftsByPatientDNI(patientDNI string)([]domain.Shift, error) */
}

type shiftRepository struct {
	storage store.StoreInterface
}

func NewRepositoryShift(storage store.StoreInterface) ShiftRepository {
	return &shiftRepository{storage}
}

func (repo *shiftRepository) GetByID(id int) (domain.Shift, error) {
	shift, err := repo.storage.ReadShift(id)
	if err != nil {
		return domain.Shift{}, errors.New("shift not found")

	}
	return shift, nil
}

func (repo *shiftRepository) CreateShift(shift domain.Shift) (domain.Shift, error) {

	return repo.storage.CreateShift(shift)

}

func (repo *shiftRepository) UpdateShift(shift domain.Shift) (domain.Shift, error) {

	return repo.storage.UpdateShift(shift)

}

func (repo *shiftRepository) DeleteShift(id int) {
	repo.storage.DeleteShift(id)
	return

}

func (r *shiftRepository) PatchShift(shift domain.Shift) (domain.Shift, error) {

	return r.storage.PatchShift(shift)

}

func (repo *shiftRepository) CreateShiftByDni(dischargeDate string, description string, dni string, registration string) (domain.Shift, error) {

	createdShift, err := repo.storage.CreateShiftByDni(dischargeDate, description, dni, registration)
	if err != nil {
		return domain.Shift{}, err
	}
	return createdShift, nil
}

/* func (r *shiftRepository) GetShiftsByPatientDNI(patientDNI string) ([]domain.Shift, error) {
	shift, err := repo.storage.ReadShift(id)
	if err != nil {
		return domain.Shift{}, errors.New("shift not found")

	}
	return shift, nil
} */
