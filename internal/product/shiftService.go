package product

import (
	"encoding/asn1"
	"repositoryapi/internal/domain"

	"github.com/go-playground/validator/v10/translations/id"
)	

type ShiftService interface {
	GetByIDShift(id int) (domain.Shift, error)
}

type shiftService struct {
	repo ShiftRepository
}

func NewServiceShift(repo ShiftRepository) ShiftService {
	return &shiftService{repo}
}

func (service *shiftService) AddShift(id int) (domain.Shift, error) {
	shift, err := service.repo.AddShift(id)
	if err != nil {
		return domain.Shift{}, err

	}
	return shift, nil

}


func (service *shiftService) GetByIDShift(id int) (domain.Shift, error) {
	shift, err := service.repo.GetByIDShift(id)
	if err != nil {
		return domain.Shift{}, err

	}
	return shift, nil
}
