package shift

import (
	"errors"
	"repositoryapi/internal/domain"
)

type ShiftServiceInterface interface {
	ReadAllShift() ([]domain.Shift, error)
	FindShiftById(id int) (domain.Shift, error)
	CreateShift(shift domain.Shift) (domain.Shift, error)
	UpdateShift(shift domain.Shift) (domain.Shift, error)
	PatchShift(shift domain.Shift) (domain.Shift, error)
	DeleteShift(id int) error
}

type shiftService struct {
	r ShiftRepositoryInterface
}

func NewShiftService(r ShiftRepositoryInterface) ShiftServiceInterface {
	return &shiftService{r}
}

func (service *shiftService) ReadAllShift() ([]domain.Shift, error) {
	shifts, err := service.r.ReadAllShift()
	if err != nil {
		return nil, err
	}
	return shifts, nil
}

func (service *shiftService) FindShiftById(id int) (domain.Shift, error) {
	shift, err := service.r.FindShiftById(id)
	if err != nil {
		return domain.Shift{}, err
	}
	return shift, nil
}

func (service *shiftService) CreateShift(shift domain.Shift) (domain.Shift, error) {
	shift, err := service.r.CreateShift(shift)
	if err != nil {
		return domain.Shift{}, err
	}
	return shift, nil
}

func (service *shiftService) UpdateShift(shift domain.Shift) (domain.Shift, error) {
	if shift.PatientID == 0 || shift.DentistID == 0 || shift.Description == "" {
		return domain.Shift{}, errors.New("all fields are required")
	}
	shift, err := service.r.UpdateShift(shift)
	if err != nil {
		return domain.Shift{}, err
	}
	return shift, nil
}

func (service *shiftService) PatchShift(shift domain.Shift) (domain.Shift, error) {
	if shift.Description == "" {
		return domain.Shift{}, errors.New("description is required")
	}
	shift, err := service.r.PatchShift(shift)
	if err != nil {
		return domain.Shift{}, err
	}
	return shift, nil
}

func (service *shiftService) DeleteShift(id int) error {
	err := service.r.DeleteShift(id)
	if err != nil {
		return err
	}
	return nil
}
