package product

import (
	"errors"
	"repositoryapi/internal/domain"
	"repositoryapi/pkg/store"
)

type DentistRepository interface {
	GetByID(id int) (domain.Dentist, error)
	/* GetAll() ([]domain.Dentist, error) */
}

type dentistRepository struct {
	storage store.StoreInterface
}

func NewRepository(storage store.StoreInterface) DentistRepository {
	return &dentistRepository{storage}
}

func (r *dentistRepository) GetByID(id int) (domain.Dentist, error) {
	dentist, err := r.storage.Read(id)
	if err != nil {
		return domain.Dentist{}, errors.New("dentist not found")

	}
	return dentist, nil
}

/*unc ( repository *dentistRepository)GetAll()([]domain.Dentist, error){

}*/
