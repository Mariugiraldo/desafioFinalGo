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
	CreateShiftByDNIAndRegistration(dni string, registration string, shift domain.Shift) (domain.Shift, error)
	ReadShiftByDNI(dni string) (domain.Shift, error)
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

func (r *shiftRepository) CreateShiftByDNIAndRegistration(dni string, registration string, shift domain.Shift) (domain.Shift, error) {
	return r.storage.CreateShiftByDNIAndRegistration(dni, registration, shift)
}

func (r *shiftRepository) ReadShiftByDNI(dni string) (domain.Shift, error) {
	return r.storage.ReadShiftByDNI(dni)
}
