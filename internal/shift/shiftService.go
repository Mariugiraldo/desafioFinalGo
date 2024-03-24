package shift

import (
	"repositoryapi/internal/domain"

)	

type ShiftService interface {
	GetByID(id int) (domain.Shift, error)
}

type shiftService struct {
	repo ShiftRepository
}

func NewServiceShift(repo ShiftRepository) ShiftService {
	return &shiftService{repo}
}

func (service *shiftService) GetByID(id int) (domain.Shift, error) {
	shift, err := service.repo.GetByID(id)
	if err != nil {
		return domain.Shift{}, err

	}
	return shift, nil
}
