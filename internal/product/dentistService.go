package product

import (
	"repositoryapi/internal/domain"
)

type DentistService interface {
	GetByID(id int) (domain.Dentist, error)
}

type dentistService struct {
	r DentistRepository
}

func NewService(r DentistRepository) DentistService {
	return &dentistService{r}
}

func (s *dentistService) GetByID(id int) (domain.Dentist, error) {
	d, err := s.r.GetByID(id)
	if err != nil {
		return domain.Dentist{}, err

	}
	return d, nil
}
