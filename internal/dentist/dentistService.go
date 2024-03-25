package dentist

import (
	"repositoryapi/internal/domain"
)

type DentistService interface {
	
	GetByID(id int) (domain.Dentist, error)
	CreateDentist(dentist domain.Dentist)(domain.Dentist, error)
	UpdateDentist(dentist domain.Dentist)(domain.Dentist, error)
	PatchDentist(dentist domain.Dentist)(domain.Dentist, error)
	Delete(id int) 
}

type dentistService struct {
	r DentistRepository
}

func NewService(r DentistRepository) DentistService {
	return &dentistService{r}
}

func (service *dentistService) GetByID(id int) (domain.Dentist, error) {
	dentist, err := service.r.GetByID(id)
	if err != nil {
		return domain.Dentist{}, err

	}
	return dentist, nil
}

func (service *dentistService) CreateDentist( dentist domain.Dentist) (domain.Dentist, error) {
    dentist, err := service.r.CreateDentist(dentist)
    if err != nil {
        return domain.Dentist{}, err
    }
    return dentist, nil
}

func (service *dentistService) UpdateDentist( dentist domain.Dentist) (domain.Dentist, error) {
    dentist, err := service.r.UpdateDentist(dentist)
    if err != nil {
        return domain.Dentist{}, err
    }
    return dentist, nil
}

func (service *dentistService) PatchDentist( dentist domain.Dentist) (domain.Dentist, error) {
    dentist, err := service.r.PatchDentist(dentist)
    if err != nil {
        return domain.Dentist{}, err
    }
    return dentist, nil
}

func (service *dentistService) Delete(id int) {
	service.r.DeleteDentist(id)

	return

	}
	
	
