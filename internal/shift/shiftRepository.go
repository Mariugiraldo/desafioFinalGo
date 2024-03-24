package shift

import (
	"errors"
	"repositoryapi/internal/domain"
	"repositoryapi/pkg/store"
)

type ShiftRepositoryInterface interface {
	ReadAllShift() ([]domain.Shift, error)
	FindShiftById(id int) (domain.Shift, error)
	CreateShift(shift domain.Shift) (domain.Shift, error)
	UpdateShift(shift domain.Shift) (domain.Shift, error)
	PatchShift(shift domain.Shift) (domain.Shift, error)
	DeleteShift(id int) error
}

type shiftRepository struct {
	storage store.ShiftStoreInterface
}

func NewShiftRepository(storage store.ShiftStoreInterface) ShiftRepositoryInterface {
	return &shiftRepository{storage}
}

func (s *shiftRepository) ReadAllShift() ([]domain.Shift, error) {
	return s.storage.ReadAllShift()
}

func (s *shiftRepository) FindShiftById(id int) (domain.Shift, error) {
	shift, err := s.storage.ReadShift(id)
	if err != nil {
		return domain.Shift{}, errors.New("Shift not found")
	}
	return shift, nil
}

func (s *shiftRepository) CreateShift(shift domain.Shift) (domain.Shift, error) {
	return s.storage.CreateShift(shift)
}

func (s *shiftRepository) UpdateShift(shift domain.Shift) (domain.Shift, error) {
	return s.storage.UpdateShift(shift)
}

func (s *shiftRepository) PatchShift(shift domain.Shift) (domain.Shift, error) {
	return s.storage.UpdateShift(shift)
}

func (s *shiftRepository) DeleteShift(id int) error {
	return s.storage.DeleteShift(id)
}
