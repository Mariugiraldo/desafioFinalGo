package product

import (
	"errors"
	"repositoryapi/internal/domain"
	"repositoryapi/pkg/store"
)

type ShiftRepository interface {
    AddShift(id int) (domain.Shift, error)
	GetByIDShift(id int) (domain.Shift, error)
}

type shiftRepository struct {
	storage store.StoreInterface
}

func NewRepositoryShift(storage store.StoreInterface) ShiftRepository {
	return &shiftRepository{storage}
}

func (repo *shiftRepository) AddShift(id int) (domain.Shift, error) {
    shift, err := repo.storage.CreateShift(id)
    if err != nil {
        return domain.Shift{}, errors.New("shift not found")
    }
    return shift, nil
}

func (repo *shiftRepository) GetByIDShift(id int) (domain.Shift, error) {
	shift, err := repo.storage.ReadShift(id)
	if err != nil {
		return domain.Shift{}, errors.New("shift not found")

	}
	return shift, nil
}