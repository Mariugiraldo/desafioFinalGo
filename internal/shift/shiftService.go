package shift

import (
	"repositoryapi/internal/domain"
)

type ShiftService interface {
	GetByID(id int) (domain.Shift, error)
	CreateShift(domain.Shift) (domain.Shift, error)
	UpdateShift(domain.Shift) (domain.Shift, error)
	DeleteShift(id int)
	PatchShift(domain.Shift) (domain.Shift, error)
	CreateShiftByDNIAndRegistration(dni string, registration string, shift domain.Shift) (domain.Shift, error)
	ReadShiftByDNI(dni string) (domain.Shift, error)
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

func (service *shiftService) CreateShift(shift domain.Shift) (domain.Shift, error) {
	shiftCreate, err := service.repo.CreateShift(shift)
	if err != nil {
		return domain.Shift{}, err
	}
	return shiftCreate, nil
}

func (service *shiftService) UpdateShift(shift domain.Shift) (domain.Shift, error) {
	shift, err := service.repo.UpdateShift(shift)
	if err != nil {
		return domain.Shift{}, err
	}
	return shift, nil
}

func (service *shiftService) DeleteShift(id int) {
	service.repo.DeleteShift(id)

	return

}

func (service *shiftService) PatchShift(shift domain.Shift) (domain.Shift, error) {
	shift, err := service.repo.PatchShift(shift)
	if err != nil {
		return domain.Shift{}, err
	}
	return shift, nil
}

func (service *shiftService) CreateShiftByDNIAndRegistration(dni string, registration string, shift domain.Shift) (domain.Shift, error) {
	shift, err := service.repo.CreateShiftByDNIAndRegistration(dni, registration, shift)
	if err != nil {
		return domain.Shift{}, err
	}
	return shift, nil
}

func (service *shiftService) ReadShiftByDNI(dni string) (domain.Shift, error) {
	shift, err := service.repo.ReadShiftByDNI(dni)
	if err != nil {
		return domain.Shift{}, err
	}
	return shift, nil
}
